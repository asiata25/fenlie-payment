package product

import "finpro-fenlie/model/entity"

type ProductRepository interface {
	GetAllProducts(page, pageSize int) ([]entity.Product, int64, error)
	InsertProduct(product entity.Product) (entity.Product, error)
	GetById(id string) (entity.Product, error)
	UpdateProduct(id string, product entity.Product) error
	SoftDeleteProduct(id string) error
}

type ProductUsecase interface {
	GetAllProducts(page, pageSize int) ([]entity.Product, int, int64, int64, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	GetProduct(id string, product entity.Product) (entity.Product, error)
	UpdateProduct(id string, product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
}
