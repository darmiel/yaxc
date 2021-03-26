package server

import (
	"errors"
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

var errEncryptionNotEnabled = errors.New("encryption not enabled")

func (s *yAxCServer) handlePostAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))

	// Read content
	bytes := ctx.Body()
	if s.MaxBodyLength > 0 && len(bytes) > s.MaxBodyLength {
		return s.errBodyLen
	}
	content := string(bytes)

	// TTL
	ttl := s.DefaultTTL
	if q := ctx.Query("ttl"); q != "" {
		if ttl, err = time.ParseDuration(q); err != nil {
			return
		}
	}

	// Encryption
	if q := ctx.Query("secret"); q != "" {
		if !s.EnableEncryption {
			return errEncryptionNotEnabled
		}
		// fail on error
		encrypt, err := common.Encrypt(content, q)
		if err != nil {
			return err
		}
		content = string(encrypt)
	}

	// Check if ttl is valid
	if (s.MinTTL != 0 && s.MinTTL > ttl) || (s.MaxTTL != 0 && s.MaxTTL < ttl) {
		return ctx.Status(400).SendString("ERROR: TTL out of range")
	}

	hash := common.Hash(content)

	// Set contents
	errVal := s.Backend.Set(path, content, ttl)
	errHsh := s.Backend.SetHash(path, hash, ttl)

	if errVal != nil || errHsh != nil {
		log.Warning("ERROR saving Value / Hash:", errVal, errHsh)
		return ctx.Status(400).SendString(
			fmt.Sprintf("ERROR (Val): %v\nERROR (Hsh): %v", errVal, errHsh))
	}

	log.Debug(ctx.IP(), "updated", path, "with hash", hash)

	return ctx.Status(200).SendString(content)
}
