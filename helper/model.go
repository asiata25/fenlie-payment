package helper

import (
	"finpro-fenlie/model/dto"
	"finpro-fenlie/model/entity"
)

func ToCompanyResponse(entity entity.Company) *dto.CompanyResponse {
	return &dto.CompanyResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		SecretKey: entity.ClientSecret,
	}
}
