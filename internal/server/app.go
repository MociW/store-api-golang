package server

import (
	"github.com/MociW/store-api-golang/internal/middleware"
	"github.com/MociW/store-api-golang/internal/user/controller"
	"github.com/MociW/store-api-golang/internal/user/repository"
	"github.com/MociW/store-api-golang/internal/user/service"
)

func (s *Server) Boostrap() error {
	middlwareSetup := middleware.NewMiddlewareManager(&middleware.MiddlewareConfig{
		Config: s.cfg,
	})

	userPostgresRepo := repository.NewUserPostgresRepository(s.db)
	userAwsRepo := repository.NewAWSUserRepository(s.awsClient)

	authService := service.NewAuthService(s.cfg, userPostgresRepo)
	userService := service.NewUserService(s.cfg, userPostgresRepo, userAwsRepo)

	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)

	user := s.app.Group("/users")
	user.Post("/", authController.RegisterNewUser)
	user.Post("/login", authController.LoginUser)

	user.Use(middlwareSetup.AuthMiddleware)
	user.Post("/avatar", userController.UploadAvatar)
	user.Get("/me", userController.GetCurrentUser)

	return nil
}
