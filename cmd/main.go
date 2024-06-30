package main

import (
	"log"
	"rest-api-example/config"
	"rest-api-example/internal/server"
	"rest-api-example/pkg/conndb"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Couldn't load config, err=%v", err)
	}

	db, err := conndb.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Couldn't load db, err=%v", err)
	}

	redis := conndb.NewRedisClient(cfg)

	s := server.NewServer(cfg, redis, db)

	if err := s.Run(); err != nil {
		log.Fatalf("Couldn't run server, err=%v", err)
	}
}
