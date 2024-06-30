package server

import (
	"rest-api-example/internal/middleware"
	"rest-api-example/internal/users/delivery/userHttp"
	userrepo "rest-api-example/internal/users/repo/postgres"
	userRedisrepo "rest-api-example/internal/users/repo/redis"
	userusecase "rest-api-example/internal/users/usecase"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(s.cfg, s.postgres)
	userRedisRepo := userRedisrepo.NewUserRepo(s.cfg, s.redis)

	mw := middleware.NewMDWManager(userRedisRepo)

	userUC := userusecase.NewUserUC(s.cfg, userRepo, userRedisRepo)
	userHandlers := userHttp.NewNewUserHandler(s.cfg, userUC)

	group := s.fiber.Group("user")
	userHttp.MapRoutes(group, userHandlers, mw)
}
