package config

import (
	"encoding/json"
	"os"
	"rest-api-example/internal/models"
)

const configPath = "./config/config.json"

type Config struct {
	Server struct {
		Host   string
		Domain string
	}
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}
	SessionSettings struct {
		SessionTTLSeconds models.TTL
	}
}

func LoadConfig() (c *Config, err error) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(jsonFile).Decode(&c)
	if err != nil {
		return nil, err
	}
	return
}
