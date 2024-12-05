package util

import (
	"bytes"
	"fmt"
	"net/http"

	product "github.com/MociW/store-api-golang/internal/product/model"
	user "github.com/MociW/store-api-golang/internal/user/model"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

func ReadUserImageRequest(c *fiber.Ctx, fieldname string) (*user.UserUploadInput, error) {
	image, err := c.FormFile(fieldname)
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
	}

	if image.Size > 1<<20 {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File size exceeds 1 MB"})
	}

	allowedContentTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/jpg":  true,
	}

	file, err := image.Open()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(file); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to read file"})
	}

	contentType := http.DetectContentType(buf.Bytes())
	if !allowedContentTypes[contentType] {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported file type"})
	}

	model := &user.UserUploadInput{
		Object:      bytes.NewReader(buf.Bytes()),
		ObjectName:  image.Filename,
		ObjectSize:  image.Size,
		BucketName:  "avatar-user-store",
		ContentType: contentType,
	}

	return model, nil

}

func ReadProductImageRequest(c *fiber.Ctx, fieldname string) (*product.ProductUploadInput, error) {
	image, err := c.FormFile(fieldname)
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
	}

	if image.Size > 3<<20 {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File size exceeds 1 MB"})
	}

	allowedContentTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/jpg":  true,
	}

	file, err := image.Open()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(file); err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to read file"})
	}

	contentType := http.DetectContentType(buf.Bytes())
	if !allowedContentTypes[contentType] {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported file type"})
	}

	fileExtension := ".jpg"
	if contentType == "image/png" {
		fileExtension = ".png"
	}

	if contentType == "image/jpeg" {
		fileExtension = ".jpeg"
	}

	newFileName := fmt.Sprintf("%s%s", GenerateShortUUID(), fileExtension)

	model := &product.ProductUploadInput{
		Object:      bytes.NewReader(buf.Bytes()),
		ObjectName:  newFileName,
		ObjectSize:  image.Size,
		BucketName:  "avatar-user-store",
		ContentType: contentType,
	}

	return model, nil

}

func GenerateShortUUID() string {
	u := uuid.New()
	return u.String()[:8]
}
