package controller

import (
	"bytes"
	"log"
	"net/http"

	"github.com/MociW/store-api-golang/internal/user"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/MociW/store-api-golang/internal/user/model/dto"
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

func (user UserControllerImpl) UpdateUser(c *fiber.Ctx) error {

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

func (user UserControllerImpl) UploadAvatar(c *fiber.Ctx) error {

	claim := c.Locals("user").(*jwt.MapClaims)

	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	image, err := c.FormFile("img")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
	}

	if image.Size > 1<<20 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File size exceeds 1 MB"})
	}

	allowedContentTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	file, err := image.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(file); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to read file"})
	}

	contentType := http.DetectContentType(buf.Bytes())
	if !allowedContentTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported file type"})
	}

	response, err := user.userService.UploadAvatar(c.UserContext(), userID, &model.UserUploadInput{
		Object:      bytes.NewReader(buf.Bytes()),
		ObjectName:  image.Filename,
		ObjectSize:  image.Size,
		BucketName:  "avatar-user-store",
		ContentType: contentType,
	})
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

func (user UserControllerImpl) GetCurrentUser(c *fiber.Ctx) error {
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
		Message: "Success Updated",
		Data:    response,
	})
}

/* --------------------------------- Address -------------------------------- */

func (user UserControllerImpl) RegisterNewAddress(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (user UserControllerImpl) UpdateAddress(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (user UserControllerImpl) FindAddress(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (user UserControllerImpl) ListAddress(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
