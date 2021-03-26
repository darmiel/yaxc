package server

import (
	"github.com/darmiel/yaxc/internal/common"
	"github.com/gofiber/fiber/v2"
)

func (s *yAxCServer) handleGetAnywhere(ctx *fiber.Ctx) (err error) {
	path := ctx.Params("anywhere")
	var res string
	if res, err = s.Backend.Get(path); err != nil {
		return
	}

	// Encryption
	if q := ctx.Query("secret"); q != "" {
		if !s.EnableEncryption {
			return errEncryptionNotEnabled
		}
		// do not fail on error
		if encrypt, err := common.Decrypt(res, q); err == nil {
			res = string(encrypt)
		}
	}

	// log.Debug(ctx.IP(), "requested", path)

	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}

func (s *yAxCServer) handleGetHashAnywhere(ctx *fiber.Ctx) (err error) {
	path := ctx.Params("anywhere")
	var res string
	if res, err = s.Backend.GetHash(path); err != nil {
		return
	}
	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}
