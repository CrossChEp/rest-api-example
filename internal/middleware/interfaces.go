package middleware

import (
	"context"
	"rest-api-example/internal/models"
)

type (
	UserRedisRepo interface {
		GetUserSession(ctx context.Context, sessionID string) (models.Claims, error)
	}
)
