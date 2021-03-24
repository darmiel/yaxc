package server

import "github.com/gofiber/fiber/v2"

func (s *yAxCServer) Start() {
	log.Info("Starting YAxC server on", s.BindAddress)
	s.App = fiber.New()

	// register routes
	s.App.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString(body)
	})

	s.App.Get("/:anywhere", s.handleGetAnywhere)
	s.App.Post("/:anywhere", s.handlePostAnywhere)

	if err := s.App.Listen(s.BindAddress); err != nil {
		log.Critical(err)
	}
}
