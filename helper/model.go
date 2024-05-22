package helper

import (
	"finpro-fenlie/model/dto/category"
	"finpro-fenlie/model/dto/company"
	"finpro-fenlie/model/entity"
)

func ToCompanyResponse(entity entity.Company) *company.CompanyResponse {
	return &company.CompanyResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		SecretKey: entity.ClientSecret,
	}
}

func ToCategoryResponse(entity entity.Category) category.CategoryResponse {
	return category.CategoryResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CompanyID: entity.CompanyID,
	}
}
