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

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with the provided details
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        data  body      dto.ProductCreateRequest  true  "Product creation request"
// @Success      201   {object}  dto.ApiProductResponse
// @Failure      400   {object}  fiber.Map
// @Failure      401   {object}  fiber.Map
// @Failure      500   {object}  fiber.Map
// @Router       /products [post]
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

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete a product by ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.ProductDeleteRequest	true	"Product delete request"
//	@Success		200		{object}	dto.ApiProductResponse
//	@Failure		400		{object}	fiber.Map
//	@Failure		500		{object}	fiber.Map
//	@Router			/products [delete]
func (product *ProductControllerImpl) DeleteProduct(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductDeleteRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	err = product.productService.DeleteProduct(c.UserContext(), request)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Delete product"})
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Product Created",
	})
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update an existing product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		dto.ProductUpdateRequest	true	"Product update request"
//	@Success		200		{object}	dto.ApiProductResponse
//	@Failure		400		{object}	fiber.Map
//	@Failure		500		{object}	fiber.Map
//	@Router			/products [put]
func (product *ProductControllerImpl) UpdateProduct(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductUpdateRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := product.productService.UpdateProduct(c.UserContext(), request)

	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Delete product"})
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Product Created",
		Data:    response,
	})
}

// FindProduct godoc
//
//	@Summary		Find a product
//	@Description	Find a product by ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.ProductFindRequest	true	"Product find request"
//	@Success		200		{object}	dto.ApiProductResponse
//	@Failure		400		{object}	fiber.Map
//	@Failure		500		{object}	fiber.Map
//	@Router			/products [get]
func (product *ProductControllerImpl) FindProduct(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductFindRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := product.productService.FindProduct(c.UserContext(), request)

	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Find Data"})
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Product Created",
		Data:    response,
	})
}

// ListProduct godoc
//
//	@Summary		List products
//	@Description	List all products for a user
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.ApiProductResponse
//	@Failure		500	{object}	fiber.Map
//	@Router			/products/list [get]
func (product *ProductControllerImpl) ListProduct(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	response, err := product.productService.ListProduct(c.UserContext(), userID)

	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Delete product"})
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Product Created",
		Data:    response,
	})
}
