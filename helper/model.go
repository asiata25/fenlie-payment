package helper

import (
	"finpro-fenlie/model/dto/category"
	"finpro-fenlie/model/dto/company"
	"finpro-fenlie/model/dto/product"
	"finpro-fenlie/model/dto/transaction"
	"finpro-fenlie/model/dto/user"
	"finpro-fenlie/model/entity"
)

func ToCompanyResponse(entity entity.Company) *company.CompanyResponse {
	return &company.CompanyResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		SecretKey: entity.SecretKey,
	}
}

func ToCategoryResponse(entity entity.Category) category.CategoryResponse {
	return category.CategoryResponse{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func ToProductResponse(entity entity.Product) product.ProductResponse {
	return product.ProductResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Price:       entity.Price,
		Description: entity.Description.String,
		Status:      entity.Status,
		Category:    entity.Category.Name,
	}
}

func ToUserResponse(entity entity.User) user.UserResponse {
	return user.UserResponse{
		ID:       entity.ID,
		Name:     entity.Name,
		Email:    entity.Email,
		Password: entity.Password,
		Role:     entity.Role,
		Company:  "",
	}
}

func ToTransactionResponse(entity entity.Transaction) transaction.ResponseTransaction {
	return transaction.ResponseTransaction{
		ID:                 entity.ID,
		OrderDate:          entity.OrderDate,
		Status:             entity.Status,
		CompanyID:          entity.CompanyID,
		UserID:             entity.UserId,
		DetailTransactions: entity.DetailTransactions,
		Invoices:           entity.Invoices,
		Total:              entity.Total,
	}
}
