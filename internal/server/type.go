package server

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/op/go-logging"
	"os"
	"time"
	"zgo.at/zcache"
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
	BindAddress string // required
	// Redis
	RedisAddress   string // "" -> only use cache
	RedisPassword  string
	RedisDatabase  int
	RedisPrefixVal string
	RedisPrefixHsh string
	// Timeout
	DefaultTTL time.Duration // 0 -> infinite
	MinTTL     time.Duration // == MaxTTL -> cannot specify TTL
	MaxTTL     time.Duration // == MinTTL -> cannot specify TTL
	// Other
	MaxBodyLength    int
	EnableEncryption bool
	ProxyHeader      string
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
			cache:   zcache.New(s.DefaultTTL, s.DefaultTTL+time.Minute),
			errCast: errors.New("not a string"),
		}
	} else {
		rb := &RedisBackend{
			ctx: context.TODO(),
			client: redis.NewClient(&redis.Options{
				Addr:     s.RedisAddress,
				Password: s.RedisPassword,
				DB:       s.RedisDatabase,
			}),
			prefixVal: s.RedisPrefixVal,
			prefixHsh: s.RedisPrefixHsh,
		}
		s.Backend = rb
		// ping test
		if cmd := rb.client.Ping(rb.ctx); cmd == nil || cmd.Err() != nil {
			log.Critical("Connection to redis failed:")
			log.Critical(cmd.Err())
			os.Exit(1)
			return
		}
	}

	if s.Backend == nil {
		log.Critical("ERROR: No backend specified.")
		os.Exit(1)
		return
	}

	log.Info("Encryption:", s.EnableEncryption)

	return
}
