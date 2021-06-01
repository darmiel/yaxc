// +build server

package server

import (
	"encoding/base64"
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/darmiel/yaxc/internal/whitelist"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/muesli/termenv"
	"net/http"
	"strings"
	"time"
)

func (s *YAxCServer) handlePostAnywhere(ctx *fiber.Ctx) (err error) {
	path := utils.CopyString(strings.TrimSpace(ctx.Params("anywhere")))
	return s.setAnywhereWithHash(ctx, path, "")
}

func (s *YAxCServer) handlePostAnywhereWithHash(ctx *fiber.Ctx) (err error) {
	path := utils.CopyString(strings.TrimSpace(ctx.Params("anywhere")))
	hash := utils.CopyString(strings.TrimSpace(ctx.Params("hash")))
	// validate hash
	if !common.ValidateHex(hash) {
		return ctx.Status(400).SendString("ERROR: Invalid hash")
	}
	return s.setAnywhereWithHash(ctx, path, hash)
}

func (s *YAxCServer) setAnywhereWithHash(ctx *fiber.Ctx, path, hash string) (err error) {
	// validate path
	if !common.ValidateAnywherePath(path) {
		return fiber.NewError(http.StatusNotAcceptable, "invalid anywhere-path")
	}

	var maxBodyLength = s.MaxBodyLength

	// check authorization
	if auth := ctx.Get("Authorization"); auth != "" {
		if len(s.JWTSign) == 0 {
			return fiber.NewError(509, "whitelist not available")
		}
		spl := strings.Split(auth, " ")
		if spl[0] != "JWT" {
			return fiber.NewError(http.StatusUnauthorized, "invalid auth type - JWT required")
		}
		claims, err := whitelist.ValidateToken(s.JWTSign, spl[1])
		if err != nil {
			return fiber.NewError(510, "error validating token: "+err.Error())
		}
		if claims.MaxBody != 0 {
			maxBodyLength = claims.MaxBody
		}
		fmt.Println(common.StyleInfo(),
			claims.Issuer, "used whitelist key with claim: mb:", claims.MaxBody, ", ia:", claims.IssuedAt)
	}

	// Read content
	bytes := ctx.Body()
	if s.MaxBodyLength > 0 && int64(len(bytes)) > maxBodyLength {
		return fiber.NewError(http.StatusRequestEntityTooLarge, "exceeded max body length")
	}
	content := string(bytes)

	// TTL
	ttl := s.DefaultTTL
	if q := ctx.Query("ttl"); q != "" {
		if ttl, err = time.ParseDuration(q); err != nil {
			return fiber.NewError(http.StatusUnprocessableEntity, "invalid ttl. (examples: 10s, 5m, 1h)")
		}
	}

	// Encryption
	if q := ctx.Query("secret"); q != "" {
		if !s.EnableEncryption {
			return fiber.NewError(http.StatusLocked, "encryption is currently not enabled on this server")
		}
		// fail on error
		encrypt, err := common.Encrypt(content, q)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, "error encrypting content: "+err.Error())
		}
		content = string(encrypt)
	}

	// Base64
	if q := ctx.Query("b64"); q != "" {
		if strings.EqualFold(q, "encode") {
			content = base64.StdEncoding.EncodeToString([]byte(content))
		} else if strings.EqualFold(q, "decode") {
			b, err := base64.StdEncoding.DecodeString(content)
			if err != nil {
				return fiber.NewError(504, "base64 encryption broke: "+err.Error())
			}
			content = string(b)
		} else {
			return fiber.NewError(506, "invalid base64 mode. available: encode, decode")
		}
	}

	// Check if ttl is valid
	if (s.MinTTL != 0 && s.MinTTL > ttl) || (s.MaxTTL != 0 && s.MaxTTL < ttl) {
		return fiber.NewError(http.StatusRequestedRangeNotSatisfiable, "ttl out of range")
	}

	// generate hash
	if hash == "" {
		hash = common.MD5Hash(content)
	}

	// Set contents
	errVal := s.Backend.Set(path, content, ttl)
	errHsh := s.Backend.SetHash(path, hash, ttl)

	if errVal != nil || errHsh != nil {
		log.Warning("ERROR saving Value / MD5Hash:", errVal, errHsh)
		return ctx.Status(500).SendString(
			fmt.Sprintf("ERROR (Val): %v\nERROR (Hsh): %v", errVal, errHsh))
	}

	fmt.Println(common.StyleServe(),
		termenv.String(ctx.IP()).Foreground(common.Profile().Color("#DBAB79")),
		"updated",
		termenv.String(path).Foreground(common.Profile().Color("#D290E4")),
		"with hash",
		termenv.String(hash).Foreground(common.Profile().Color("#71BEF2")))

	return ctx.Status(200).SendString(content)
}
