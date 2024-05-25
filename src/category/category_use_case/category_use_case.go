package categoryUseCase

import (
	"finpro-fenlie/helper"
	categoryDto "finpro-fenlie/model/dto/category"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/category"
)

type categoryUseCase struct {
	repository category.CategoryRepository
}

// Create implements category.CategoryUseCase.
func (uc *categoryUseCase) Create(request *categoryDto.CategoryRequest) error {
	category := entity.Category{
		Name:      request.Name,
		CompanyID: request.CompanyID,
	}

	err := uc.repository.Save(&category)

	return err
}

// GetAlls implements category.CategoryUseCase.
func (uc *categoryUseCase) GetAll(page, size int, name, companyId string) (*[]categoryDto.CategoryResponse, int, error) {
	var categories []categoryDto.CategoryResponse
	results, total, err := uc.repository.GetAll(page, size, name, companyId)
	if err != nil {
		return nil, 0, err
	}

	for _, res := range *results {
		categories = append(categories, helper.ToCategoryResponse(res))
	}

	return &categories, int(total), nil
}

// GetById implements category.CategoryUseCase.
func (uc *categoryUseCase) GetById(ID, companyId string) (categoryDto.CategoryResponse, error) {
	category, err := uc.repository.GetById(ID, companyId)
	if err != nil {
		return categoryDto.CategoryResponse{}, err
	}

	return helper.ToCategoryResponse(category), nil
}

// Update implements category.CategoryUseCase.
func (uc *categoryUseCase) Update(request *categoryDto.CategoryRequest) error {
	category := entity.Category{
		ID:        request.ID,
		Name:      request.Name,
		CompanyID: request.CompanyID,
	}

	err := uc.repository.Update(&category)
	return err
}

// Delete implements category.CategoryUseCase
func (uc *categoryUseCase) Delete(ID, companyId string) error {

	err := uc.repository.Delete(ID, companyId)
	return err
}

func NewCategoryUseCase(repository category.CategoryRepository) category.CategoryUseCase {
	return &categoryUseCase{repository}
}
