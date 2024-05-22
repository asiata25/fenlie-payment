package companyUseCase

import (
	"finpro-fenlie/helper"
	companyDTO "finpro-fenlie/model/dto/company"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/company"

	"golang.org/x/crypto/bcrypt"
)

type companyUseCase struct {
	repository company.CompanyRepository
}

// GetAll implements company.CompanyUseCase.
func (c *companyUseCase) GetAll() ([]*companyDTO.CompanyResponse, error) {
	var companies []*companyDTO.CompanyResponse

	results, err := c.repository.FindAll()
	if err != nil {
		return companies, err
	}

	for _, result := range results {
		companies = append(companies, helper.ToCompanyResponse(*result))
	}

	return companies, nil
}

// Create implements company.CompanyUseCase.
func (c *companyUseCase) Create(request companyDTO.CompanyCreateRequest) error {
	company := entity.Company{
		Name:         request.Name,
		Email:        request.Email,
		ClientSecret: request.SecretKey,
	}

	hashedSecret, err := bcrypt.GenerateFromPassword([]byte(company.ClientSecret), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	company.ClientSecret = string(hashedSecret)

	error := c.repository.Save(company)
	return error
}

// Delete implements company.CompanyUseCase.
func (c *companyUseCase) Delete(id string) error {
	err := c.repository.Delete(id)
	return err
}

// GetById implements company.CompanyUseCase.
func (c *companyUseCase) GetById(id string) (*companyDTO.CompanyResponse, error) {
	company, err := c.repository.RetrieveByID(id)
	if err = helper.CheckErrNotFound(err); err != nil {
		return &companyDTO.CompanyResponse{}, err
	}

	return helper.ToCompanyResponse(*company), nil
}

// Update implements company.CompanyUseCase.
func (c *companyUseCase) Update(request companyDTO.CompanyUpdateRequest) error {
	companyExisting, err := c.repository.RetrieveByID(request.ID)
	if err != nil {
		return err
	}

	if request.SecretKey == "" {
		request.SecretKey = companyExisting.ClientSecret
	} else {
		hashedSecret, err := bcrypt.GenerateFromPassword([]byte(request.SecretKey), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		request.SecretKey = string(hashedSecret)
	}

	company := entity.Company{
		ID:           request.ID,
		Name:         request.Name,
		ClientSecret: request.SecretKey,
	}

	err = c.repository.Update(company)
	if err = helper.CheckErrNotFound(err); err != nil {
		return err
	}

	return nil
}

func NewCompanyUseCase(repository company.CompanyRepository) company.CompanyUseCase {
	return &companyUseCase{repository}
}
