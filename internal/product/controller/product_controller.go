package controller

import (
	"encoding/json"
	"log"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/MociW/store-api-golang/internal/product/model/dto"
	"github.com/MociW/store-api-golang/pkg/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ProductControllerImpl struct {
	productService product.ProductService
}

func NewProductContoller(productService product.ProductService) product.ProductController {
	return &ProductControllerImpl{productService: productService}
}

func (product *ProductControllerImpl) CreateProduct(c *fiber.Ctx) error {
	// Extract user ID from the token
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	// Parse the request body
	request := new(dto.ProductCreateRequest)
	data := c.FormValue("data")
	if err := json.Unmarshal([]byte(data), &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON data"})
	}

	// Assign user ID to the request
	request.UserID = userID

	// Initialize a slice for storing product image inputs
	var objs []model.ProductUploadInput

	// Process product image uploads
	products := []string{"image_01", "image_02", "image_03"}
	for _, product := range products {
		obj, err := util.ReadProductImageRequest(c, product)
		if err != nil {
			log.Printf("Error reading image for %s: %v", product, err)
			continue // Skip this product if there's an error
		}
		if obj != nil { // Add only non-nil objects to the slice
			objs = append(objs, *obj)
		}
	}

	// Assign the images to the request
	request.Images = objs

	// Call the service to create the product
	response, err := product.productService.CreateProduct(c.UserContext(), request)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	// Return the response
	return c.Status(fiber.StatusCreated).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusCreated,
		Message: "Product Created",
		Data:    response,
	})
}
