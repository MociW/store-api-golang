package user

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	RegisterNewUser(c *fiber.Ctx) error

	ValidateUser(c *fiber.Ctx) error

	LoginUser(c *fiber.Ctx) error
}

type UserController interface {
	/* ---------------------------------- User ---------------------------------- */

	UpdateUser(c *fiber.Ctx) error

	UploadAvatar(c *fiber.Ctx) error

	GetCurrentUser(c *fiber.Ctx) error

	/* --------------------------------- Address -------------------------------- */

	RegisterNewAddress(c *fiber.Ctx) error

	UpdateAddress(c *fiber.Ctx) error

	FindAddress(c *fiber.Ctx) error

	ListAddress(c *fiber.Ctx) error
}
