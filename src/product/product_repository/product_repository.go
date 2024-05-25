package productRepository

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/product"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	return &productRepository{db}
}

func (repo *productRepository) GetAllProducts(page, pageSize int, name, companyId string) ([]entity.Product, int, error) {
	var products []entity.Product
	var totalItems int64

	if err := repo.db.Model(&entity.Product{}).Scopes(helper.FindBasedOnCompany(companyId), helper.Paginate(page, pageSize)).Where("products.name LIKE $1", "%"+name+"%").Count(&totalItems).Joins("Category", repo.db.Select("Category.name")).Error; err != nil {
		return nil, 0, err
	}

	return products, int(totalItems), nil
}

func (repo *productRepository) InsertProduct(product entity.Product) error {
	if err := repo.db.Create(&product).Error; err != nil {
		return err
	}
	return nil

}

func (repo *productRepository) GetById(id, companyId string) (entity.Product, error) {
	var product entity.Product
	if err := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Joins("Category", repo.db.Select("Category.name")).First(&product, "products.id = ?", id).Error; err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (repo *productRepository) UpdateProduct(product entity.Product) error {
	// TODO: Updates using map for allowing zero value
	err := repo.db.Model(&product).Scopes(helper.FindBasedOnCompany(product.CompanyID)).Omit("id", "company_id").Updates(product).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *productRepository) DeleteProduct(id, companyId string) error {
	result := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Delete(&entity.Product{ID: id})
	if result.RowsAffected < 1 {
		return errors.New("cannot find the requested data")
	}

	return nil
}
