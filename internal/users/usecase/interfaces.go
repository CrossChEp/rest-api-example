package usecase

import (
	"context"
	"rest-api-example/internal/models"
)

type (
	UserRepo interface {
		Create(ctx context.Context, user models.User) (models.UserID, error)
		Get(ctx context.Context, filter models.UserFilter) (models.User, error)
		GetMany(ctx context.Context, filter models.UserFilter) ([]models.User, error)
	}

	UserRedisRepo interface {
		SetUserSession(ctx context.Context, sessionID string, claims models.Claims) error
	}
)
