package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type MDWManger struct {
	userRedisRepo UserRedisRepo
}

func NewMDWManager(userRedisRepo UserRedisRepo) *MDWManger {
	return &MDWManger{
		userRedisRepo: userRedisRepo,
	}
}

func (m *MDWManger) AuthedMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Cookies("session_key")

		if sessionKey == "" {
			c.ClearCookie("session_key")
		}

		claims, err := m.userRedisRepo.GetUserSession(c.Context(), sessionKey)
		if err != nil {
			c.ClearCookie("session_key")
			return err
		}

		if claims.UserID == 0 {
			c.ClearCookie("session_key")
			return errors.New("unauthorized")
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
