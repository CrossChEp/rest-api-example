package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"strings"
	"time"
)

type MDWManager struct {
	cfg      *config.Config
	userRepo UserRepo
}

func NewMDWManager(cfg *config.Config, userRepo UserRepo) *MDWManager {
	return &MDWManager{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

// AuthedMiddleware
// 1. Извлечь токен из хедера
// 2. Bearer d,dasifofkewow
func (m *MDWManager) AuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()["Authorization"][0]

		if header == "" {
			return errors.New("auth header is empty")
		}

		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			return errors.New("invalid header value")
		}

		token := parts[1]
		claims := &models.Claims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return m.cfg.PublicKey, nil
		})
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if claims.ExpiresAt.Before(time.Now()) && c.Path() != "/user/refresh" {
			return errors.New("unauthorized")
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
