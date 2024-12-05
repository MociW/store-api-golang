package server

import (
	"github.com/MociW/store-api-golang/internal/middleware"
	productController "github.com/MociW/store-api-golang/internal/product/controller"
	productRepo "github.com/MociW/store-api-golang/internal/product/repository"
	productService "github.com/MociW/store-api-golang/internal/product/service"
	userController "github.com/MociW/store-api-golang/internal/user/controller"
	userRepo "github.com/MociW/store-api-golang/internal/user/repository"
	userService "github.com/MociW/store-api-golang/internal/user/service"
)

func (s *Server) Boostrap() error {
	middlewareSetup := middleware.NewMiddlewareManager(&middleware.MiddlewareConfig{
		Config: s.cfg,
	})

	/* ----------------------------- User Repository ---------------------------- */

	userPostgresRepo := userRepo.NewUserPostgresRepository(s.db)
	userAwsRepo := userRepo.NewAWSUserRepository(s.awsClient)

	/* --------------------------- Product Repository --------------------------- */

	productAWSRepo := productRepo.NewProductAWSRepository(s.awsClient)
	productRepo := productRepo.NewProductRepository(s.db)

	/* ------------------------------ User Service ------------------------------ */

	authService := userService.NewAuthService(s.cfg, userPostgresRepo)
	userService := userService.NewUserService(s.cfg, userPostgresRepo, userAwsRepo)

	/* ----------------------------- Product Service ---------------------------- */

	productService := productService.NewProductService(s.cfg, productRepo, productAWSRepo)

	/* ----------------------------- User Controller ---------------------------- */

	authController := userController.NewAuthController(authService)
	userController := userController.NewUserController(userService)

	/* --------------------------- Product Controller --------------------------- */

	productController := productController.NewProductContoller(productService)

	user := s.app.Group("/users")
	user.Post("/", authController.RegisterNewUser)
	user.Post("/login", authController.LoginUser)

	user.Use(middlewareSetup.AuthMiddleware)
	user.Post("/avatar", userController.UploadAvatar)
	user.Get("/me", userController.GetCurrentUser)
	user.Post("/addresses", userController.RegisterNewAddress)

	product := s.app.Group("/products")
	product.Use(middlewareSetup.AuthMiddleware)
	product.Use("/add-product", productController.CreateProduct)

	return nil
}
