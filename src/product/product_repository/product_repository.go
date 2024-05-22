package productRepository

import (
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/product"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	return &productRepository{db}
}

func (repo *productRepository) GetAllProducts(page, pageSize int) ([]entity.Product, int64, error) {
	var products []entity.Product
	var totalItems int64

	offset := (page - 1) * pageSize
	if err := repo.db.Model(&entity.Product{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	if err := repo.db.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalItems, nil
}

func (repo *productRepository) InsertProduct(product entity.Product) (entity.Product, error) {
	if err := repo.db.Create(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil

}

func (repo *productRepository) GetById(id string) (entity.Product, error) {
	var product entity.Product
	if err := repo.db.First(&product, "id = ?", id).Error; err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (repo *productRepository) UpdateProduct(id string, product entity.Product) error {
	err := repo.db.Model(&product).Where("id = ?", id).Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *productRepository) SoftDeleteProduct(id string) error {
	var product entity.Product

	if err := repo.db.Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
