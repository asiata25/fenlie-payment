package product

import (
	"finpro-fenlie/model/dto/product"
	"finpro-fenlie/model/entity"
)

type ProductRepository interface {
	GetAllProducts(page, pageSize int, name, companyId string) ([]entity.Product, int, error)
	InsertProduct(product entity.Product) error
	GetById(id, companyId string) (entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(id, companyId string) error
}

type ProductUsecase interface {
	GetAllProducts(page, pageSize int, name, companyId string) ([]product.ProductResponse, int, error)
	CreateProduct(request product.ProductCreateRequest, imageURL string) error
	GetProduct(id, companyId string) (product.ProductResponse, error)
	UpdateProduct(request product.ProductUpdateRequest) error
	DeleteProduct(id, companyId string) error
}
