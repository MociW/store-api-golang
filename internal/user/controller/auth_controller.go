package controller

import (
	"time"

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

func (auth *AuthControllerImpl) RegisterNewUser(c *fiber.Ctx) error {
	request := new(dto.UserRegisterRequest)
	err := c.BodyParser(request)
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := auth.authService.Register(c.UserContext(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ApiUserResponse{
		Status:  fiber.StatusCreated,
		Message: "Account Created Successfully",
		Data:    response,
	})
}

func (auth *AuthControllerImpl) LoginUser(c *fiber.Ctx) error {
	request := new(dto.UserLoginRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	response, err := auth.authService.Login(c.UserContext(), request)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	const accessTokenExpiry = 15 * time.Minute
	const refreshTokenExpiry = 24 * time.Hour

	// Set cookies
	c.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    response.AccessToken,
		Expires:  time.Now().Add(accessTokenExpiry),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    response.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(dto.ApiUserResponse{
		Status:  fiber.StatusOK,
		Message: "Login successful",
		Data:    response,
	})
}
