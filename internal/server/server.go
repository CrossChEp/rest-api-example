package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"os"
	"os/signal"
	"rest-api-example/config"
	"syscall"
)

type Server struct {
	cfg      *config.Config
	postgres *sqlx.DB
	fiber    *fiber.App
	redis    *redis.Client
}

func NewServer(cfg *config.Config, redis *redis.Client, postgres *sqlx.DB) *Server {
	return &Server{
		cfg:      cfg,
		postgres: postgres,
		redis:    redis,
		fiber:    fiber.New(),
	}
}

func (s *Server) Run() error {
	s.MapHandlers()

	go func() {
		if err := s.fiber.Listen(s.cfg.Server.Host); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err := s.fiber.Shutdown()
	if err != nil {
		fmt.Println("Server finished with panic")
	} else {
		fmt.Println("HTTP server closed properly")
	}

	return nil
}
