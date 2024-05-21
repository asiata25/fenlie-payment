package router

import (
	companyDelivery "finpro-fenlie/src/company/company_delivery"
	companyRepository "finpro-fenlie/src/company/company_repository"
	companyUseCase "finpro-fenlie/src/company/company_use_case"
	"finpro-fenlie/src/user/userDelivery"
	"finpro-fenlie/src/user/userRepository"
	"finpro-fenlie/src/user/userUsecase"

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
}
