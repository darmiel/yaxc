package server

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/op/go-logging"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

var (
	log    = logging.MustGetLogger("example")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

func init() {
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend2Formatter)
}

type YAxCConfig struct {
	// Address
	BindAddress  string // required
	RedisAddress string // "" -> only use cache
	// Timeout
	DefaultTTL time.Duration // 0 -> infinite
	MinTTL     time.Duration // == MaxTTL -> cannot specify TTL
	MaxTTL     time.Duration // == MinTTL -> cannot specify TTL
	// Other
	MaxBodyLength int
}

type yAxCServer struct {
	*YAxCConfig
	App        *fiber.App
	Backend    Backend
	errBodyLen error
}

func NewServer(cfg *YAxCConfig) (s *yAxCServer) {
	s = &yAxCServer{
		YAxCConfig: cfg,
		errBodyLen: errors.New("exceeded max body length"),
	}

	// backend
	if s.RedisAddress == "" {
		// use cache backend
		s.Backend = &CacheBackend{
			c:       cache.New(s.DefaultTTL, s.DefaultTTL+time.Minute),
			errCast: errors.New("not a string"),
		}
	} else {
		// TODO: implement redis
	}

	if s.Backend == nil {
		log.Critical("ERROR: No backend specified.")
		os.Exit(1)
		return
	}

	return
}

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
