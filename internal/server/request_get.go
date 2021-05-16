package server

import (
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/gofiber/fiber/v2"
	"github.com/muesli/termenv"
	"net/http"
	"strings"
)

func (s *YAxCServer) handleGetAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))

	// validate path
	if !common.ValidateAnywherePath(path) {
		return fiber.NewError(http.StatusNotAcceptable, "invalid anywhere-path")
	}

	var res string
	if res, err = s.Backend.Get(path); err != nil {
		return
	}

	// Encryption
	if q := ctx.Query("secret"); q != "" {
		if !s.EnableEncryption {
			return fiber.NewError(http.StatusLocked, "encryption is currently not enabled on this server")
		}
		// do not fail on error
		if encrypt, err := common.Decrypt(res, q); err == nil {
			res = string(encrypt)
		}
	}

	fmt.Println(common.StyleServe(),
		termenv.String(ctx.IP()).Foreground(common.Profile().Color("#DBAB79")),
		"requested",
		termenv.String("value").Foreground(common.Profile().Color("#A8CC8C")),
		termenv.String(path).Foreground(common.Profile().Color("#D290E4")))

	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}

func (s *YAxCServer) handleGetHashAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	var res string
	if res, err = s.Backend.GetHash(path); err != nil {
		return
	}

	fmt.Println(common.StyleServe(),
		termenv.String(ctx.IP()).Foreground(common.Profile().Color("#DBAB79")),
		"requested",
		termenv.String("hash").Foreground(common.Profile().Color("#E88388")),
		termenv.String(path).Foreground(common.Profile().Color("#D290E4")))

	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}
