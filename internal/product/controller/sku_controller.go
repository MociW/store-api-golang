package controller

import (
	"log"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ProductSKUContollerImpl struct {
	SKUService product.ProductSKUService
}

func NewProductSKUController(skuService product.ProductSKUService) product.ProductSKUContoller {
	return &ProductSKUContollerImpl{SKUService: skuService}
}

// CreateSKU godoc
//	@Summary		Create a SKU
//	@Description	Create a new SKU for a product
//	@Tags			sku
//	@Accept			json
//	@Produce		json
//	@Param			sku	body		dto.ProductSKUCreateRequest	true	"SKU data"
//	@Success		200	{object}	dto.ApiProductResponse
//	@Failure		400	{object}	fiber.Map
//	@Failure		500	{object}	fiber.Map
//	@Router			/sku [post]
func (product *ProductSKUContollerImpl) CreateSKU(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductSKUCreateRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := product.SKUService.CreateSKU(c.UserContext(), request)
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

// DeleteSKU godoc
//	@Summary		Delete a SKU
//	@Description	Delete a SKU by ID
//	@Tags			sku
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.ProductSKUDeleteRequest	true	"SKU delete request"
//	@Success		200		{object}	dto.ApiProductResponse
//	@Failure		400		{object}	fiber.Map
//	@Failure		500		{object}	fiber.Map
//	@Router			/sku [delete]
func (product *ProductSKUContollerImpl) DeleteSKU(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductSKUDeleteRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	err = product.SKUService.DeleteSKU(c.UserContext(), request)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Delete product"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Product Deleted",
	})
}

func (product *ProductSKUContollerImpl) FindSKU(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductSKUFindRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := product.SKUService.FindSKU(c.UserContext(), request)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Find Data"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Product Found",
		Data:    response,
	})
}

func (product *ProductSKUContollerImpl) ListSKU(c *fiber.Ctx) error {
	claim := c.Locals("user").(*jwt.MapClaims)
	userID, ok := (*claim)["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	request := new(dto.ProductSKUListRequest)
	err := c.BodyParser(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	request.UserID = userID

	response, err := product.SKUService.ListSKU(c.UserContext(), request.ProductID, request.UserID)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to Find Data"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ApiProductResponse{
		Status:  fiber.StatusOK,
		Message: "Products Found",
		Data:    response,
	})
}
