package controller

import (
	"log"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
	"github.com/MociW/store-api-golang/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserControllerImpl struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) user.UserController {
	return &UserControllerImpl{userService: userService}
}

/* ---------------------------------- User ---------------------------------- */

func (user *UserControllerImpl) UpdateUser(c *fiber.Ctx) error {

	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.UserUpdateRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := user.userService.UpdateUser(c.UserContext(), request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to Update Data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiUserResponse{
		Status:  fiber.StatusOK,
		Message: "Success Updated",
		Data:    response,
	})
}

func (user *UserControllerImpl) UploadAvatar(c *fiber.Ctx) error {

	claim := c.Locals("user").(*jwt.MapClaims)

	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	image, err := util.ReadUserImageRequest(c, "img")
	if err != nil {
		log.Printf("Error uploading avatar: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve image"})
	}

	response, err := user.userService.UploadAvatar(c.UserContext(), userID, image)
	if err != nil {
		log.Printf("Error uploading avatar: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload avatar"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiUserResponse{
		Status:  fiber.StatusOK,
		Message: "Avatar uploaded successfully",
		Data:    map[string]any{"avatar": response.Avatar},
	})
}

func (user *UserControllerImpl) GetCurrentUser(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)

	email, ok := (*claim)["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	response, err := user.userService.GetCurrentUser(c.UserContext(), email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiUserResponse{
		Status:  fiber.StatusOK,
		Message: "Success Retrieving Data",
		Data:    response,
	})
}

/* --------------------------------- Address -------------------------------- */

func (user *UserControllerImpl) RegisterNewAddress(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)

	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.CreateAddressRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := user.userService.CreateAddress(c.UserContext(), request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Failed to Create Data",
		})
	}

	return c.Status(201).JSON(dto.ApiUserResponse{
		Status:  201,
		Message: "Successfully Create Data",
		Data:    response,
	})
}

func (user *UserControllerImpl) UpdateAddress(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)

	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.UpdateAddressRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := user.userService.UpdateAddress(c.UserContext(), request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Failed to Update Data",
		})
	}

	return c.Status(200).JSON(dto.ApiUserResponse{
		Status:  200,
		Message: "Successfully Update Data",
		Data:    response,
	})
}

func (user *UserControllerImpl) FindAddress(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)

	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.FindAddressRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := user.userService.FindAddress(c.UserContext(), request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Failed to Retreive Data",
		})
	}

	return c.Status(200).JSON(dto.ApiUserResponse{
		Status:  200,
		Message: "Successfully Retreive Data",
		Data:    response,
	})

}

func (user *UserControllerImpl) ListAddress(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)

	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	response, err := user.userService.ListAddress(c.UserContext(), userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ApiUserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Failed to Retreive Data",
		})
	}

	return c.Status(200).JSON(dto.ApiUserResponse{
		Status:  200,
		Message: "Successfully Retreive Data",
		Data:    response,
	})
}
