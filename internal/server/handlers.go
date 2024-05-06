package server

import (
	userHttp "rest-api-example/internal/users/delivery/http"
	userrepo "rest-api-example/internal/users/repo/postgres"
	userrredis "rest-api-example/internal/users/repo/redis"
	userusecase "rest-api-example/internal/users/usecase"
)

func (s *Server) MapHandlers() {
	userRepo := userrepo.NewUserRepo(s.cfg, s.postgres)
	userRedisRepo := userrredis.NewUserRedisRepo(s.redis, s.cfg)
	userUC := userusecase.NewUserUC(s.cfg, userRepo, userRedisRepo)
	userHTTPHandlers := userHttp.NewUserHandlers(s.cfg, userUC)

	userGroup := s.fiber.Group("user")
	userHttp.MapUserRoutes(userGroup, userHTTPHandlers)
}
