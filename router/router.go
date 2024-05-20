package router

import (
	companyDelivery "finpro-fenlie/src/company/company_delivery"
	companyRepository "finpro-fenlie/src/company/company_repository"
	companyUseCase "finpro-fenlie/src/company/company_use_case"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(v1Group *gin.RouterGroup, db *gorm.DB) {
	companyRepository := companyRepository.NewCompanyRepository(db)
	companyUseCase := companyUseCase.NewCompanyUseCase(companyRepository)
	companyDelivery.NewCompanyDelivery(v1Group, companyUseCase)
}
