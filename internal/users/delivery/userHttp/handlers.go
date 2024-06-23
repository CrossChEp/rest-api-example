package userHttp

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	userusecase "rest-api-example/internal/users/usecase"
)

type UserHandlers struct {
	userUC UserUC
}

func NewNewUserHandler(userUC UserUC) *UserHandlers {
	return &UserHandlers{
		userUC: userUC,
	}
}

func (h *UserHandlers) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := RegisterRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		uid, err := h.userUC.Register(c.Context(), userusecase.RegisterUser{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": fmt.Sprintf("User was registered with id=%d", uid),
		})
	}
}
