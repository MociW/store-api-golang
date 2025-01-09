package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/MociW/store-api-golang/internal/product"
	"github.com/MociW/store-api-golang/internal/product/model"
	"github.com/MociW/store-api-golang/internal/product/model/dto"
	"github.com/MociW/store-api-golang/pkg/config"
)

type ProductServiceImpl struct {
	cfg         *config.Config
	productRepo product.ProductRepository
	awsRepo     product.ProductAWSRepository
}

func NewProductService(cfg *config.Config, productRepo product.ProductRepository, awsRepo product.ProductAWSRepository) product.ProductService {
	return &ProductServiceImpl{cfg: cfg, productRepo: productRepo, awsRepo: awsRepo}
}

func (product *ProductServiceImpl) UploadImage(ctx context.Context, id string, file *model.ProductUploadInput) (string, error) {
	uploadInfo, err := product.awsRepo.PutObject(ctx, file)
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	if file.BucketName == "" || product.cfg.AWS.Endpoint == "" {
		return "", fmt.Errorf("invalid bucket name or endpoint")
	}

	avatarURL := generateAWSProductURL(file.BucketName, uploadInfo.Key, product.cfg.AWS.Endpoint)
	return avatarURL, nil
}

func generateAWSProductURL(bucket string, key string, endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint, bucket, key)
}

// func extractAWSProductURL(url string) string {
// 	return
// }

func (product *ProductServiceImpl) CreateProduct(ctx context.Context, entity *dto.ProductCreateRequest) (*dto.ProductResponse, error) {
	// Validate input
	if entity == nil {
		return nil, fmt.Errorf("invalid product request: entity is nil")
	}

	var wg sync.WaitGroup
	mu := sync.Mutex{} // Protect shared resources
	images := []string{}
	errChan := make(chan error, len(entity.Images))

	// Concurrent uploads
	for _, image := range entity.Images {
		wg.Add(1)
		go func(img model.ProductUploadInput) {
			defer wg.Done()
			uploadedURL, err := product.UploadImage(ctx, entity.UserID, &img)
			if err != nil {
				errChan <- fmt.Errorf("failed to upload image %s: %w", img.ObjectName, err)
				return
			}
			mu.Lock()
			images = append(images, uploadedURL)
			mu.Unlock()
		}(image)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errChan)

	// Log errors (if any)
	for err := range errChan {
		// Replace this with a proper logging mechanism if necessary
		fmt.Println(err)
	}

	// Prepare request
	request := &model.Product{
		Name:        entity.Name,
		Description: entity.Description,
		Summary:     entity.Summary,
		UserID:      entity.UserID,
		Images:      images,
	}

	// Save product
	result, err := product.productRepo.CreateProduct(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	// Convert to response DTO
	return dto.ConvertProductResponse(result), nil
}

func (product *ProductServiceImpl) UpdateProduct(ctx context.Context, entity *dto.ProductUpdateRequest) (*dto.ProductResponse, error) {

	request := &model.Product{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Summary:     entity.Summary,
		UserID:      entity.UserID,
	}

	result, err := product.productRepo.UpdateProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertProductResponse(result), nil
}

func (product *ProductServiceImpl) DeleteProduct(ctx context.Context, entity *dto.ProductDeleteRequest) error {
	request := &model.Product{
		ID:     entity.ID,
		UserID: entity.UserID,
	}

	// result, err := product.productRepo.FindProduct(ctx, request)
	// if err != nil {
	// 	return err
	// }

	err := product.productRepo.DeleteProduct(ctx, request)
	if err != nil {
		return err
	}

	// for _, image := range result.Images {
	// 	err := product.awsRepo.RemoveObject(ctx, "product-store", path.Base(image))
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (product *ProductServiceImpl) FindProduct(ctx context.Context, entity *dto.ProductFindRequest) (*dto.ProductResponse, error) {
	request := &model.Product{
		ID:     entity.ID,
		UserID: entity.UserID,
	}

	result, err := product.productRepo.FindProduct(ctx, request)
	if err != nil {
		return nil, err
	}

	return dto.ConvertProductResponse(result), nil
}

func (product *ProductServiceImpl) ListProduct(ctx context.Context, id string) ([]dto.ProductResponse, error) {
	result, err := product.productRepo.ListProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ProductResponse, len(result))
	for i, address := range result {
		responses[i] = *dto.ConvertProductResponse(&address)
	}

	return responses, nil
}

// func (product *ProductSKUServiceImpl) generateListProductPdf(ctx context.Context, id string) error {
// 	tableHeadings := []string{"Product", "SKU", "Price", "Quantity"}
// 	content:=
// }
