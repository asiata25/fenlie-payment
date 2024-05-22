package productUsecase

import (
	"errors"
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

func (uc *productUC) GetAllProducts(page, pageSize int) ([]entity.Product, int, int64, int64, error) {
	products, totalItems, err := uc.productRepo.GetAllProducts(page, pageSize)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize) // Calculate total pages
	return products, page, totalPages, totalItems, nil
}

func (uc *productUC) CreateProduct(product entity.Product) (entity.Product, error) {
	result, err := uc.productRepo.InsertProduct(product)
	if err != nil {
		return entity.Product{}, err
	}

	return result, nil
}

func (uc *productUC) GetProduct(id string, product entity.Product) (entity.Product, error) {
	result, err := uc.productRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, errors.New("product not found")
		}
		return entity.Product{}, err
	}

	return result, nil
}

func (uc *productUC) UpdateProduct(id string, product entity.Product) (entity.Product, error) {
	existingProduct, err := uc.productRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, errors.New("product not found")
		}
		return entity.Product{}, err
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Description = product.Description
	existingProduct.Status = product.Status
	existingProduct.CategoryID = product.CategoryID
	existingProduct.CompanyID = product.CompanyID

	err = uc.productRepo.UpdateProduct(id, existingProduct)
	if err != nil {
		return entity.Product{}, err
	}

	return existingProduct, nil
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
