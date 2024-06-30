package userHttp

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rest-api-example/internal/models"
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

func (h *UserHandlers) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := LoginRequest{}
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		token, err := h.userUC.SignIn(c.Context(), userusecase.SignIn{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": token,
		})
	}
}

func (h *UserHandlers) GetOwn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		user, err := h.userUC.GetUser(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": user,
		})
	}
}

func (h *UserHandlers) RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*models.Claims)
		if !ok {
			return errors.New("cannot get claims")
		}

		token, err := h.userUC.RefreshToken(c.Context(), claims.UserID)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"data": token,
		})
	}
}
