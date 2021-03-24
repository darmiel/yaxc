package server

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func (s *yAxCServer) handlePostAnywhere(ctx *fiber.Ctx) (err error) {
	path := ctx.Params("anywhere")

	bytes := ctx.Body()
	if s.MaxBodyLength > 0 && len(bytes) > s.MaxBodyLength {
		return s.errBodyLen
	}
	content := string(bytes)

	// custom ttl
	ttl := s.DefaultTTL

	if q := ctx.Query("ttl"); q != "" {
		if ttl, err = time.ParseDuration(q); err != nil {
			return
		}
	}

	// check if ttl is valid
	if (s.MinTTL != 0 && s.MinTTL > ttl) || (s.MaxTTL != 0 && s.MaxTTL < ttl) {
		return ctx.Status(400).SendString("ERROR: TTL out of range")
	}

	// set contents
	if err := s.Backend.Set(path, content, ttl); err != nil {
		return ctx.Status(400).SendString("ERROR: " + err.Error())
	}

	log.Info("Received", content, "to", path, "with a ttl of", ttl)
	return ctx.Status(200).SendString(content)
}

func (s *yAxCServer) handleGetAnywhere(ctx *fiber.Ctx) (err error) {
	path := ctx.Params("anywhere")
	var res string
	if res, err = s.Backend.Get(path); err != nil {
		return
	}
	log.Info("Send", res, "to", ctx.IP(), "requesting", path)

	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}
