package http

import "github.com/gofiber/fiber/v2"

func MapUserRoutes(group fiber.Router, h Handlers) {
	group.Post("/register", h.Register())
	group.Post("/sign_in", h.SignIn())
}
