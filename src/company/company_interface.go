package company

import (
	"finpro-fenlie/model/dto"
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
	Create(request dto.CompanyCreateRequest) error
	Update(request dto.CompanyUpdateRequest) error
	Delete(id string) error
	GetById(id string) (*dto.CompanyResponse, error)
	GetAll() ([]*dto.CompanyResponse, error)
}
