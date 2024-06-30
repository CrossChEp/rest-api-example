package userHttp

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	userusecase "rest-api-example/internal/users/usecase"
	"time"
)

type UserHandlers struct {
	cfg    *config.Config
	userUC UserUC
}

func NewNewUserHandler(cfg *config.Config, userUC UserUC) *UserHandlers {
	return &UserHandlers{
		cfg:    cfg,
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

		sessionKey, err := h.userUC.SignIn(c.Context(), userusecase.SignIn{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session_key",
			Value:    sessionKey,
			Path:     "/",
			Secure:   true,
			HTTPOnly: true,
			Domain:   h.cfg.Server.Domain,
			Expires:  time.Now().Add(time.Second * time.Duration(h.cfg.SessionSettings.SessionTTLSeconds)),
		},
		)

		return c.JSON(fiber.Map{
			"data": "user signed in successfully",
		})
	}
}

func (h *UserHandlers) GetOwn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(models.Claims)
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
