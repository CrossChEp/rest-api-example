package userHttp

import (
	"github.com/gofiber/fiber/v2"
	"rest-api-example/internal/middleware"
)

func MapRoutes(group fiber.Router, h Handlers, mw *middleware.MDWManager) {
	group.Post("/register", h.Register())
	group.Post("/login", h.Login())
	group.Post("/get_own", mw.AuthedMiddleware(), h.GetOwn())
	group.Post("/refresh", mw.AuthedMiddleware(), h.RefreshToken())
}
