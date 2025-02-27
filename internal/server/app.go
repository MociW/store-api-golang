package server

import (
	_ "github.com/MociW/store-api-golang/docs"
	"github.com/MociW/store-api-golang/internal/middleware"
	productController "github.com/MociW/store-api-golang/internal/product/controller"
	productRepo "github.com/MociW/store-api-golang/internal/product/repository"
	productService "github.com/MociW/store-api-golang/internal/product/service"
	userController "github.com/MociW/store-api-golang/internal/user/controller"
	userRepo "github.com/MociW/store-api-golang/internal/user/repository"
	userService "github.com/MociW/store-api-golang/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
)

func (s *Server) Bootstrap() error {
	middlewareSetup := middleware.NewMiddlewareManager(&middleware.MiddlewareConfig{
		Config: s.cfg,
		Logger: s.logger,
	})

	/* ----------------------------- User Repository ---------------------------- */

	UserPostgresRepo := userRepo.NewUserPostgresRepository(s.db)
	UserAwsRepo := userRepo.NewAWSUserRepository(s.awsClient)

	/* --------------------------- Product Repository --------------------------- */

	ProductAWSRepo := productRepo.NewProductAWSRepository(s.awsClient)
	ProductRepo := productRepo.NewProductRepository(s.db)
	SkuRepo := productRepo.NewProductSKURepository(s.db)

	/* ------------------------------ User Service ------------------------------ */

	AuthService := userService.NewAuthService(s.cfg, UserPostgresRepo, s.redis, s.mail)
	UserService := userService.NewUserService(s.cfg, UserPostgresRepo, UserAwsRepo)

	/* ----------------------------- Product Service ---------------------------- */

	ProductService := productService.NewProductService(s.cfg, ProductRepo, ProductAWSRepo)
	SkuService := productService.NewProductSKUService(s.cfg, SkuRepo)

	/* ----------------------------- User Controller ---------------------------- */

	AuthController := userController.NewAuthController(AuthService)
	UserController := userController.NewUserController(UserService)

	/* --------------------------- Product Controller --------------------------- */

	ProductController := productController.NewProductController(ProductService)
	SkuController := productController.NewProductSKUController(SkuService)

	s.app.Use(middlewareSetup.LoggerMidddleware)

	app := s.app.Group("/api/v1")
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Pong!")
	})
	app.Get("/swagger/*", swagger.HandlerDefault)

	user := app.Group("/users")
	user.Post("/", AuthController.RegisterNewUser)
	user.Post("/validate-otp", AuthController.ValidateUser)
	user.Post("/login", AuthController.LoginUser)

	user.Use(middlewareSetup.AuthMiddleware)
	user.Get("/me", UserController.GetCurrentUser)
	user.Post("/me/avatar", UserController.UploadAvatar)

	user.Get("/me/addresses", UserController.ListAddress)
	user.Get("/me/addresses/:address_id", UserController.FindAddress)
	user.Post("/me/addresses", UserController.RegisterNewAddress)

	product := app.Group("/products")
	product.Use(middlewareSetup.AuthMiddleware)
	product.Get("/", ProductController.ListProduct)
	product.Get("/:id", ProductController.FindProduct)
	product.Post("/", ProductController.CreateProduct)
	product.Delete("/:id", ProductController.DeleteProduct)
	product.Put("/:id", ProductController.UpdateProduct)

	product.Get("/:id", SkuController.ListSKU)
	product.Get("/:id/skus/:sku_id", SkuController.FindSKU)
	product.Post("/:id/skus", SkuController.CreateSKU)
	product.Delete("/:id/skus/:sku_id", SkuController.DeleteSKU)

	return nil
}
