package categoryUseCase

import (
	"finpro-fenlie/helper"
	categoryDto "finpro-fenlie/model/dto/category"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/category"
	"strconv"
)

type categoryUseCase struct {
	repository category.CategoryRepository
}

// CreateLoan implements category.CategoryUseCase.
func (uc *categoryUseCase) CreateLoan(request *categoryDto.CategoryRequest) error {

	category := entity.Category{
		Name: request.Name,
	}

	err := uc.repository.Save(&category)
	if err != nil {
		return err
	}

	return nil
}

// GetAllLoans implements category.CategoryUseCase.
func (uc *categoryUseCase) GetAllLoans(page, size string) (*[]categoryDto.CategoryResponse, int, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, 0, err
	}

	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return nil, 0, err
	}

	if sizeInt == 0 {
		sizeInt = 10
	}

	var categories []categoryDto.CategoryResponse
	results, total, err := uc.repository.GetAll(pageInt, sizeInt)
	if err != nil {
		return nil, 0, err
	}

	for _, res := range *results {
		categories = append(categories, helper.ToCategoryResponse(res))
	}

	return &categories, int(total), nil
}

// GetById implements category.CategoryUseCase.
func (uc *categoryUseCase) GetLoanById(ID string) (categoryDto.CategoryResponse, error) {
	category, err := uc.repository.GetById(ID)
	if err != nil {
		return categoryDto.CategoryResponse{}, err
	}

	return helper.ToCategoryResponse(category), nil
}

// Update implements category.CategoryUseCase.
func (uc *categoryUseCase) UpdateLoan(request *categoryDto.CategoryRequest) error {

	category := entity.Category{
		Name: request.Name,
	}

	err := uc.repository.Update(&category)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements category.CategoryUseCase
func (uc *categoryUseCase) DeleteLoan(ID string) error {

	err := uc.repository.Delete(ID)
	return err
}

func NewCategoryUseCase(repository category.CategoryRepository) category.CategoryUseCase {
	return &categoryUseCase{repository}
}
