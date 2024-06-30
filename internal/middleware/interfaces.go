package middleware

import (
	"context"
	"rest-api-example/internal/models"
)

type (
	UserRepo interface {
		Get(ctx context.Context, filter models.UserFilter) (models.User, error)
	}
)
