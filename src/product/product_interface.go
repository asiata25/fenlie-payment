package product

import (
	"finpro-fenlie/model/dto/product"
	"finpro-fenlie/model/entity"
)

type ProductRepository interface {
	GetAllProducts(page, pageSize int) ([]entity.Product, int64, error)
	InsertProduct(product entity.Product) (entity.Product, error)
	GetById(id string) (entity.Product, error)
	UpdateProduct(id string, product entity.Product) error
	SoftDeleteProduct(id string) error
}

type ProductUsecase interface {
	GetAllProducts(page, pageSize int) ([]product.ProductResponse, int, int64, int64, error)
	CreateProduct(product product.ProductRequest) error
	GetProduct(id string) (product.ProductResponse, error)
	UpdateProduct(id string, product product.ProductRequest) error
	DeleteProduct(id string) error
}
