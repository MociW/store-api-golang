package product

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	CreateProduct(c *fiber.Ctx) error

	DeleteProduct(c *fiber.Ctx) error

	FindProduct(c *fiber.Ctx) error

	UpdateProduct(c *fiber.Ctx) error
	ListProduct(c *fiber.Ctx) error
}

type ProductSKUContoller interface {
	CreateSKU(c *fiber.Ctx) error

	DeleteSKU(c *fiber.Ctx) error

	FindSKU(c *fiber.Ctx) error

	ListSKU(c *fiber.Ctx) error
}
