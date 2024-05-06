package config

import (
	"encoding/json"
	"os"
	"rest-api-example/internal/models"
)

const configPath = "./config/config.json"

type Config struct {
	Server struct {
		Host         string
		CookieDomain string
	}
	Postgres struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
	Cache struct {
		Session struct {
			Host               string
			Port               string
			MinIdleConns       int
			PoolSize           int
			PoolTimeout        int
			Password           string
			UseCertificates    bool
			InsecureSkipVerify bool
			CertificatesPaths  struct {
				Cert string
				Key  string
				Ca   string
			}
			DB int
		}
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
