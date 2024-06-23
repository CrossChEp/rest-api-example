package userHttp

import "github.com/gofiber/fiber/v2"

func MapRoutes(group fiber.Router, h Handlers) {
	group.Post("/register", h.Register())
}
