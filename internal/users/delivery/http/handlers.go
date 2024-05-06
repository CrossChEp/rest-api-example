package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"rest-api-example/config"
	"time"
)

type UserHandlers struct {
	cfg    *config.Config
	userUC UserUC
}

func NewUserHandlers(cfg *config.Config, userUC UserUC) *UserHandlers {
	return &UserHandlers{
		cfg:    cfg,
		userUC: userUC,
	}
}

func (h *UserHandlers) Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		req := RegisterRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		registerUserDTO := req.toRegisterUser()

		userID, err := h.userUC.Register(ctx.Context(), registerUserDTO)
		if err != nil {
			return err
		}

		return ctx.JSON(fiber.Map{
			"data": fmt.Sprintf("registered userID=%d", userID),
		})
	}
}

func (h *UserHandlers) SignIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := SignInRequest{}
		if err := ctx.BodyParser(&req); err != nil {
			return err
		}

		signInDTO := req.toSignIn()

		session, err := h.userUC.SignIn(ctx.Context(), signInDTO)
		if err != nil {
			return err
		}

		ctx.Cookie(&fiber.Cookie{
			Name:    "session_key",
			Value:   session.SessionKey,
			Expires: time.Now().Add(time.Second * time.Duration(h.cfg.SessionSettings.SessionTTLSeconds)),
		})

		return ctx.JSON(fiber.Map{
			"data": "user was signed in",
		})
	}
}
