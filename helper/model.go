package helper

import (
	"finpro-fenlie/model/dto/categoryDto"
	"finpro-fenlie/model/entity"
)

func ToCategoryResponse(entity entity.Category) categoryDto.CategoryResponse {
	return categoryDto.CategoryResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CompanyID: entity.CompanyID,
	}
}
