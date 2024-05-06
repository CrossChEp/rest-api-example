package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"rest-api-example/config"
	"syscall"
)

type Server struct {
	cfg      *config.Config
	fiber    *fiber.App
	postgres *sqlx.DB
	redis    *redis.Client
}

func NewServer(cfg *config.Config, postgres *sqlx.DB, redis *redis.Client) *Server {
	return &Server{
		cfg: cfg,
		fiber: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
		postgres: postgres,
		redis:    redis,
	}
}

func (s *Server) Run() error {
	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins:     "https://chainik.fun, https://chainik.pro",
		AllowCredentials: true,
	}))

	s.MapHandlers()

	go func() {
		if err := s.fiber.Listen(s.cfg.Server.Host); err != nil {
			log.Fatalf("Couldn't start fiber server, err=%v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := s.fiber.Shutdown(); err != nil {
		return err
	}
	return nil
}
