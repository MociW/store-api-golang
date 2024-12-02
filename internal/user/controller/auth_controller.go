package controller

import (
	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	authService user.AuthService
}

func NewAuthController(authService user.AuthService) user.AuthController {
	return &AuthControllerImpl{authService: authService}
}

func (auth AuthControllerImpl) RegisterNewUser(c *fiber.Ctx) error {
	request := new(dto.UserRegisterRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := auth.authService.Register(c.UserContext(), request)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(response)
}

func (auth AuthControllerImpl) LoginUser(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
