package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"rest-api-example/config"
	"syscall"
)

type Server struct {
	cfg      *config.Config
	postgres *sqlx.DB
	fiber    *fiber.App
}

func NewServer(cfg *config.Config, postgres *sqlx.DB) *Server {
	return &Server{
		cfg:      cfg,
		postgres: postgres,
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
