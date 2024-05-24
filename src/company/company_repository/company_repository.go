package companyRepository

import (
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/company"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type companyRepository struct {
	db *gorm.DB
}

// FindAll implements company.CompanyRepository.
func (c *companyRepository) FindAll() ([]*entity.Company, error) {
	var companies []*entity.Company
	err := c.db.Find(&companies).Error
	if err != nil {
		return companies, err
	}

	return companies, nil
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
func (c *companyRepository) RetrieveByID(id string) (*entity.Company, error) {
	var company entity.Company

	result := c.db.Where("id = $1", id).Take(&company)
	if result.RowsAffected < 1 {
		return &company, errors.New("cannot find the requested data")
	}
	return &company, nil
}

// Save implements company.CompanyRepository.
func (c *companyRepository) Save(payload entity.Company) error {
	err := c.db.Create(&payload).Error
	return err
}

// Update implements company.CompanyRepository.
func (c *companyRepository) Update(payload entity.Company) error {
	fmt.Println("SINI BANK", payload)
	err := c.db.Debug().Model(&payload).Omit("email").Updates(map[string]interface{}{
		"name":       payload.Name,
		"secret_key": payload.SecretKey,
	}).Error
	return err
}

func NewCompanyRepository(db *gorm.DB) company.CompanyRepository {
	return &companyRepository{db}
}
