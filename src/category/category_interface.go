package category

import (
	categoryDto "finpro-fenlie/model/dto/category"
	"finpro-fenlie/model/entity"
)

type CategoryRepository interface {
	Save(category *entity.Category) error
	GetAll(page, size int, name, companyId string) (*[]entity.Category, int64, error)
	GetById(id, companyId string) (entity.Category, error)
	Update(category *entity.Category) error
	Delete(id, companyId string) error
}

type CategoryUseCase interface {
	Create(request *categoryDto.CategoryRequest) error
	GetAll(page, size int, name, companyId string) (*[]categoryDto.CategoryResponse, int, error)
	GetById(id, companyId string) (categoryDto.CategoryResponse, error)
	Update(request *categoryDto.CategoryRequest) error
	Delete(id, companyId string) error
}
