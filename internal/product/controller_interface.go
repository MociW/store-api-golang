package product

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	CreateProduct(c *fiber.Ctx) error
}

type ProductSKUContoller interface {
}
