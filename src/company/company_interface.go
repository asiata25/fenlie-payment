package company

import (
	"finpro-fenlie/model/dto/company"
	"finpro-fenlie/model/entity"
)

type CompanyRepository interface {
	Save(payload entity.Company) error
	Update(payload entity.Company) error
	Delete(id string) error
	RetrieveByID(id string) (*entity.Company, error)
	FindAll() ([]*entity.Company, error)
}

type CompanyUseCase interface {
	Create(request company.CompanyCreateRequest) error
	Update(request company.CompanyUpdateRequest) error
	Delete(id string) error
	GetById(id string) (*company.CompanyResponse, error)
	GetAll() ([]*company.CompanyResponse, error)
}
