package controller

import (
	"errors"
	"time"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
	"github.com/MociW/store-api-golang/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthControllerImpl struct {
	authService user.AuthService
}

func NewAuthController(authService user.AuthService) user.AuthController {
	return &AuthControllerImpl{authService: authService}
}

// RegisterNewUser  godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.UserRegisterRequest	true	"User  registration request"
//	@Success		201		{object}	dto.ApiUserResponse
//	@Failure		400		{object}	dto.ApiUserResponse
//	@Failure		500		{object}	dto.ApiUserResponse
//	@Failure		409		{object}	dto.ApiUserResponse
//	@Router			/users [post]
func (auth *AuthControllerImpl) RegisterNewUser(c *fiber.Ctx) error {
	request := new(dto.UserRegisterRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
		})
	}

	if err := validator.ValidateStruct(c.UserContext(), request); err != nil {
		result := validator.TranslateValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Validation failed",
			Data:    result,
		})
	}

	response, err := auth.authService.Register(c.UserContext(), request)
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return c.Status(fiber.StatusConflict).JSON(dto.ApiUserResponse{
				Status:  fiber.StatusConflict,
				Message: "User already exists",
				Data:    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ApiUserResponse{
		Status:  fiber.StatusCreated,
		Message: "Account Created Successfully",
		Data:    response,
	})
}

// LoginUser  godoc
//
//	@Summary		User login
//	@Description	Authenticate a user and return access and refresh tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.UserLoginRequest	true	"User  login request"
//	@Success		200		{object}	dto.ApiUserResponse
//	@Failure		400		{object}	dto.ApiUserResponse
//	@Failure		401		{object}	dto.ApiUserResponse
//	@Failure		500		{object}	dto.ApiUserResponse
//	@Router			/users/login [post]
func (auth *AuthControllerImpl) LoginUser(c *fiber.Ctx) error {
	request := new(dto.UserLoginRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := validator.ValidateStruct(c.UserContext(), request); err != nil {
		result := validator.TranslateValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Validation failed",
			Data:    result,
		})
	}

	response, err := auth.authService.Login(c.UserContext(), request)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ApiUserResponse{
				Status:  fiber.StatusNotFound,
				Message: "User not found",
			})
		}

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiUserResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid Password",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
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
