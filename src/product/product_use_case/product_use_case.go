package productUsecase

import (
	"database/sql"
	"finpro-fenlie/helper"
	productDTO "finpro-fenlie/model/dto/product"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/product"
)

type productUC struct {
	productRepo product.ProductRepository
}

func NewProductUsecase(productRepo product.ProductRepository) product.ProductUsecase {
	return &productUC{productRepo: productRepo}
}

func (uc *productUC) GetAllProducts(page, pageSize int, name, companyId string) ([]productDTO.ProductResponse, int, error) {
	results, totalItems, err := uc.productRepo.GetAllProducts(page, pageSize, name, companyId)
	if err != nil {
		return nil, 0, err
	}

	var products []productDTO.ProductResponse
	for _, result := range results {
		products = append(products, helper.ToProductResponse(result))
	}

	return products, totalItems, nil
}

func (uc *productUC) CreateProduct(request productDTO.ProductCreateRequest, imageURL string) error {

	product := entity.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: sql.NullString{String: request.Description},
		Status:      request.Status,
		CategoryID:  sql.NullString{String: request.CategoryID},
		CompanyID:   request.CompanyID,
		Image:       imageURL,
	}

	error := uc.productRepo.InsertProduct(product)
	return error
}

func (uc *productUC) GetProduct(id, companyId string) (productDTO.ProductResponse, error) {
	result, err := uc.productRepo.GetById(id, companyId)
	if err != nil {
		return productDTO.ProductResponse{}, err
	}

	product := helper.ToProductResponse(result)
	product.Category = result.Category.Name

	return product, nil
}

func (uc *productUC) UpdateProduct(request productDTO.ProductUpdateRequest) error {
	product := entity.Product{
		ID:          request.ID,
		Name:        request.Name,
		Price:       request.Price,
		Description: sql.NullString{String: request.Description, Valid: request.Description != ""},
		Status:      request.Status,
		CategoryID:  sql.NullString{String: request.CategoryID, Valid: request.CategoryID != ""},
		CompanyID:   request.CompanyID,
	}

	err := uc.productRepo.UpdateProduct(product)
	return err
}

func (uc *productUC) DeleteProduct(id, companyId string) error {
	if err := uc.productRepo.DeleteProduct(id, companyId); err != nil {
		return err
	}

	return nil
}
