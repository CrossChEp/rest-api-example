package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"os"
	"rest-api-example/internal/models"
)

const (
	configPath = "./config/config.json"
	keyPath    = "./config/ec_key.pem"
)

type Config struct {
	Server struct {
		Host string
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

	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
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

	pem, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	c.PrivateKey, err = jwt.ParseECPrivateKeyFromPEM(pem)
	if err != nil {
		return
	}

	c.PublicKey = &c.PrivateKey.PublicKey
	return
}
