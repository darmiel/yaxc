package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"strings"
	"time"
)

func (s *yAxCServer) Start() {
	log.Info("Starting YAxC server on", s.BindAddress)

	cfg := &fiber.Config{}
	if s.ProxyHeader != "" {
		if s.ProxyHeader == "$proxy" {
			s.ProxyHeader = "X-Forwarded-For"
		}
		cfg.ProxyHeader = s.ProxyHeader
	}
	s.App = fiber.New(*cfg)

	// limiter middleware
	s.App.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			if c.IP() == "127.0.0.1" {
				return true
			}
			if strings.HasPrefix(c.Path(), "/hash/") {
				return true
			}
			return false
		},
		Max:        65,
		Expiration: 60 * time.Second,
	}))

	// register routes
	s.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(body)
	})

	// GET contents
	s.App.Get("/:anywhere", s.handleGetAnywhere)

	// GET hash
	s.App.Get("/hash/:anywhere", s.handleGetHashAnywhere)

	// SET contents, auto hash
	s.App.Post("/:anywhere", s.handlePostAnywhere)

	// SET contents, custom hash
	s.App.Post("/:anywhere/:hash", s.handlePostAnywhereWithHash)

	if err := s.App.Listen(s.BindAddress); err != nil {
		log.Critical(err)
		panic(err)
	}
}
