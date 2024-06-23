package server

import (
	"rest-api-example/internal/users/delivery/userHttp"
	userrepo "rest-api-example/internal/users/repo/postgres"
	userusecase "rest-api-example/internal/users/usecase"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(s.cfg, s.postgres)
	userUC := userusecase.NewUserUC(s.cfg, userRepo)
	userHandlers := userHttp.NewNewUserHandler(userUC)

	group := s.fiber.Group("user")
	userHttp.MapRoutes(group, userHandlers)
}
