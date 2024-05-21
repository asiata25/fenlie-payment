package category

import (
	categoryDto "finpro-fenlie/model/dto/category"
	"finpro-fenlie/model/entity"
)

type CategoryRepository interface {
	Save(loan *entity.Category) error
	GetAll(page, size int) (*[]entity.Category, int64, error)
	GetById(id string) (entity.Category, error)
	Update(loan *entity.Category) error
	Delete(id string) error
}

type CategoryUseCase interface {
	CreateLoan(request *categoryDto.CreateCategoryRequest) error
	GetAllLoans(page, size string) (*[]categoryDto.CategoryResponse, int, error)
	GetLoanById(ID string) (categoryDto.CategoryResponse, error)
	UpdateLoan(request *categoryDto.UpdateCategoryRequest) error
	DeleteLoan(ID string) error
}
