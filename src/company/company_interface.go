package company

import (
	"finpro-fenlie/model/dto/company"
	"finpro-fenlie/model/entity"
)

type CompanyRepository interface {
	Save(payload entity.Company) (string, error)
	Update(payload entity.Company) error
	Delete(id string) error
	RetrieveByID(id string) (entity.Company, error)
	FindAll(page, size int, name string) ([]entity.Company, int64, error)
}

type CompanyUseCase interface {
	Create(request company.CompanyCreateRequest) error
	Update(request company.CompanyUpdateRequest) error
	Delete(id string) error
	GetById(id string) (company.CompanyResponse, error)
	GetAll(page, size int, name string) ([]company.CompanyResponse, int64, error)
}
