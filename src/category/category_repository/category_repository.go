package categoryRepository

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/category"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

// Save implements category.categoryRepository.
func (repo *categoryRepository) Save(category *entity.Category) error {
	err := repo.db.Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements category.categoryRepository.
func (repo *categoryRepository) GetAll(page, size int, name, companyId string) (*[]entity.Category, int64, error) {
	var categories []entity.Category
	var total int64

	err := repo.db.Model(&entity.Category{}).Scopes(helper.FindBasedOnCompany(companyId), helper.Paginate(page, size)).Where("name LIKE $1", "%"+name+"%").Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	err = repo.db.Model(&entity.Category{}).Where("name LIKE $1", "%"+name+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return &categories, total, nil
}

// GetById implements category.categoryRepository.
func (repo *categoryRepository) GetById(id, companyId string) (entity.Category, error) {
	var category entity.Category
	result := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Preload("Products").Where("id = $1", id).Take(&category)
	if result.Error != nil {
		return category, result.Error
	}

	return category, nil
}

// Update implements category.categoryRepository.
func (repo *categoryRepository) Update(category *entity.Category) error {
	err := repo.db.Model(category).Scopes(helper.FindBasedOnCompany(category.CompanyID)).Omit("id", "company_id").Where("id = ?", category.ID).Updates(category).Error
	return err
}

// Delete implements category.categoryRepository.
func (repo *categoryRepository) Delete(id, companyId string) error {
	result := repo.db.Scopes(helper.FindBasedOnCompany(companyId)).Delete(&entity.Category{ID: id})
	if result.RowsAffected < 1 {
		return errors.New("cannot find the requested data")
	}

	return nil
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return &categoryRepository{db}
}
