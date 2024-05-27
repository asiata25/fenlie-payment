package companyRepository

import (
	"finpro-fenlie/helper"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/company"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type companyRepository struct {
	db *gorm.DB
}

// FindAll implements company.CompanyRepository.
func (c *companyRepository) FindAll(page, size int, name string) ([]entity.Company, int64, error) {
	var companies []entity.Company
	var total int64

	err := c.db.Model(&entity.Company{}).Scopes(helper.Paginate(page, size)).Where("name LIKE $1", "%"+name+"%").Preload("Users").Find(&companies).Error
	if err != nil {
		return nil, 0, err
	}

	err = c.db.Model(&entity.Company{}).Where("name LIKE $1", "%"+name+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return companies, total, nil
}

// Delete implements company.CompanyRepository.
func (c *companyRepository) Delete(id string) error {
	result := c.db.Select(clause.Associations).Delete(&entity.Company{ID: id})

	if result.RowsAffected < 1 {
		return errors.New("cannot find the requested data")
	}
	return nil
}

// RetrieveByID implements company.CompanyRepository.
func (c *companyRepository) RetrieveByID(id string) (entity.Company, error) {
	var company entity.Company

	result := c.db.Where("id = ?", id).Take(&company)
	if result.RowsAffected < 1 {
		return company, errors.New("cannot find the requested data")
	}
	return company, nil
}

// Save implements company.CompanyRepository.
func (c *companyRepository) Save(payload entity.Company) (string, error) {
	tx := c.db.Begin()
	if tx.Error != nil {
		return "", tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create Company
	err := tx.Omit("Users").Create(&payload).Error
	if err != nil {
		tx.Rollback()
		return "", errors.New(err.Error())
	}

	err = tx.Create(&entity.User{
		Name:      payload.Users[0].Name,
		Email:     payload.Users[0].Email,
		Password:  payload.Users[0].Password,
		Role:      payload.Users[0].Role,
		CompanyID: payload.ID,
	}).Error
	if err != nil {
		tx.Rollback()
		return "", errors.New(err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}
	return payload.ID, nil
}

// Update implements company.CompanyRepository.
func (c *companyRepository) Update(payload entity.Company) error {
	err := c.db.Debug().Model(&payload).Omit("email").Updates(map[string]interface{}{
		"name":       payload.Name,
		"secret_key": payload.SecretKey,
	}).Error
	return err
}

func NewCompanyRepository(db *gorm.DB) company.CompanyRepository {
	return &companyRepository{db}
}
