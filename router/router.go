package router

import (
	"finpro-fenlie/src/category/categoryDelivery"
	"finpro-fenlie/src/category/categoryRepository"
	"finpro-fenlie/src/category/categoryUseCase"
	companyDelivery "finpro-fenlie/src/company/company_delivery"
	companyRepository "finpro-fenlie/src/company/company_repository"
	companyUseCase "finpro-fenlie/src/company/company_use_case"
	userDelivery "finpro-fenlie/src/user/user_delivery"
	userRepository "finpro-fenlie/src/user/user_repository"
	userUsecase "finpro-fenlie/src/user/user_use_case"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(v1Group *gin.RouterGroup, db *gorm.DB) {

	// Company Route
	companyRepository := companyRepository.NewCompanyRepository(db)
	companyUseCase := companyUseCase.NewCompanyUseCase(companyRepository)
	companyDelivery.NewCompanyDelivery(v1Group, companyUseCase)

	// User Route
	userRepo := userRepository.NewUserRepository(db)
	userUC := userUsecase.NewUserUsecase(userRepo)
	userDelivery.NewUserDelivery(v1Group, userUC)

	// Category Route
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryDelivery.NewCategoryDelivery(v1Group, categoryUseCase)

}
