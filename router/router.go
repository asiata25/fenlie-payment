package router

import (
	categoryDelivery "finpro-fenlie/src/category/category_delivery"
	categoryRepository "finpro-fenlie/src/category/category_repository"
	categoryUseCase "finpro-fenlie/src/category/category_use_case"
	companyDelivery "finpro-fenlie/src/company/company_delivery"
	companyRepository "finpro-fenlie/src/company/company_repository"
	companyUseCase "finpro-fenlie/src/company/company_use_case"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(v1Group *gin.RouterGroup, db *gorm.DB) {

	// Company Route
	companyRepository := companyRepository.NewCompanyRepository(db)
	companyUseCase := companyUseCase.NewCompanyUseCase(companyRepository)
	companyDelivery.NewCompanyDelivery(v1Group, companyUseCase)

	// Category Route
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryDelivery.NewCategoryDelivery(v1Group, categoryUseCase)

}
