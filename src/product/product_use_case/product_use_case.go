package productUsecase

import (
	"errors"
	"finpro-fenlie/helper"
	productDTO "finpro-fenlie/model/dto/product"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/product"

	"gorm.io/gorm"
)

type productUC struct {
	productRepo product.ProductRepository
}

func NewProductUsecase(productRepo product.ProductRepository) product.ProductUsecase {
	return &productUC{productRepo: productRepo}
}

func (uc *productUC) GetAllProducts(page, pageSize int) ([]productDTO.ProductResponse, int, int64, int64, error) {
	results, totalItems, err := uc.productRepo.GetAllProducts(page, pageSize)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	var products []productDTO.ProductResponse
	for _, result := range results {
		products = append(products, helper.ToProductResponse(result))
	}

	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize) // Calculate total pages
	return products, page, totalPages, totalItems, nil
}

func (uc *productUC) CreateProduct(request productDTO.ProductCreateRequest) error {
	product := entity.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Status:      request.Status,
		CategoryID:  request.CategoryID,
		CompanyID:   request.CompanyID,
	}

	error := uc.productRepo.InsertProduct(product)
	return error
}

func (uc *productUC) GetProduct(id string) (productDTO.ProductResponse, error) {
	result, err := uc.productRepo.GetById(id)
	if err != nil {
		return productDTO.ProductResponse{}, err
	}

	return helper.ToProductResponse(result), nil
}

func (uc *productUC) UpdateProduct(id string, product productDTO.ProductRequest) error {
	existingProduct, err := uc.productRepo.GetById(id)
	if err != nil {
		return err
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Description = product.Description
	existingProduct.Status = product.Status
	existingProduct.CategoryID = product.CategoryID

	err = uc.productRepo.UpdateProduct(id, existingProduct)
	if err != nil {
		return err
	}

	return nil
}

func (uc *productUC) DeleteProduct(id string) error {
	_, err := uc.productRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if err := uc.productRepo.SoftDeleteProduct(id); err != nil {
		return err
	}

	return nil
}
