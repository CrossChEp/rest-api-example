package userHttp

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"rest-api-example/internal/models"
	"rest-api-example/internal/users/usecase"
)

type (
	UserUC interface {
		Register(ctx context.Context, regData usecase.RegisterUser) (models.UserID, error)
		SignIn(ctx context.Context, signInData usecase.SignIn) (string, error)
		GetUser(ctx context.Context, userID models.UserID) (models.User, error)
		RefreshToken(ctx context.Context, userID models.UserID) (string, error)
	}

	Handlers interface {
		Register() fiber.Handler
		Login() fiber.Handler
		GetOwn() fiber.Handler
		RefreshToken() fiber.Handler
	}
)
