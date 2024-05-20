package categoryRepository

import (
	"errors"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/category"
	"time"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

// Save implements category.categoryRepository.
func (repo *categoryRepository) Save(category *entity.Category) error {
	err := repo.db.Debug().Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements category.categoryRepository.
func (repo *categoryRepository) GetAll(page, size int) (*[]entity.Category, int64, error) {
	var categories []entity.Category
	var total int64

	offset := (page - 1) * size

	err := repo.db.Model(&entity.Category{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = repo.db.Where("deleted_at IS NULL").
		Limit(size).
		Offset(offset).
		Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	return &categories, total, nil
}

// GetById implements category.categoryRepository.
func (repo *categoryRepository) GetById(id string) (entity.Category, error) {
	var category entity.Category
	result := repo.db.Where("id = $1", id).Take(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return category, errors.New("category not found")

		}
		return category, result.Error
	}

	return category, nil
}

// Update implements category.categoryRepository.
func (repo *categoryRepository) Update(category *entity.Category) error {
	err := repo.db.Model(category).Where("id = ?", category.ID).Updates(category).Error
	return err
}

// Delete implements category.categoryRepository.
func (repo *categoryRepository) Delete(id string) error {
	currentTime := time.Now()
	err := repo.db.Model(&entity.Category{}).Where("id = ?", id).Update("DeletedAt", currentTime).Error

	return err
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return &categoryRepository{db}
}
