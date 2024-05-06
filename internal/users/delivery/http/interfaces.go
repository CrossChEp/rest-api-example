package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"rest-api-example/internal/models"
	"rest-api-example/internal/users/usecase"
)

type (
	UserUC interface {
		Register(ctx context.Context, regData usecase.RegisterUser) (models.UserID, error)
		SignIn(ctx context.Context, signInData usecase.SignIn) (models.Session, error)
	}

	Handlers interface {
		Register() fiber.Handler
		SignIn() fiber.Handler
	}
)
