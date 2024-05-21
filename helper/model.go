package helper

import (
	"finpro-fenlie/model/dto"
	"finpro-fenlie/model/dto/categoryDto"
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

func ToCategoryResponse(entity entity.Category) categoryDto.CategoryResponse {
	return categoryDto.CategoryResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CompanyID: entity.CompanyID,
	}
}
