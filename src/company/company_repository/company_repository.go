package companyRepository

import (
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/company"
	"fmt"

	"gorm.io/gorm"
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
	err := c.db.Delete(&entity.Company{ID: id}).Error

	return err
}

// RetrieveByID implements company.CompanyRepository.
func (c *companyRepository) RetrieveByID(id string) (*entity.Company, error) {
	var company entity.Company

	err := c.db.Where("id = $1", id).Take(&company).Error
	if err != nil {
		return &company, err
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
		"name":          payload.Name,
		"client_secret": payload.ClientSecret,
	}).Error
	return err
}

func NewCompanyRepository(db *gorm.DB) company.CompanyRepository {
	return &companyRepository{db}
}
