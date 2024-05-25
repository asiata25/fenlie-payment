package router

import (
	"finpro-fenlie/pkg/middleware"
	categoryDelivery "finpro-fenlie/src/category/category_delivery"
	categoryRepository "finpro-fenlie/src/category/category_repository"
	categoryUseCase "finpro-fenlie/src/category/category_use_case"
	companyDelivery "finpro-fenlie/src/company/company_delivery"
	companyRepository "finpro-fenlie/src/company/company_repository"
	companyUseCase "finpro-fenlie/src/company/company_use_case"
	productDelivery "finpro-fenlie/src/product/product_delivery"
	productRepository "finpro-fenlie/src/product/product_repository"
	productUseCase "finpro-fenlie/src/product/product_use_case"
	transactionDelivery "finpro-fenlie/src/transaction/transaction_delivery"
	transactionRepository "finpro-fenlie/src/transaction/transaction_repository"
	transactionUseCase "finpro-fenlie/src/transaction/transaction_use_case"
	userDelivery "finpro-fenlie/src/user/user_delivery"
	userRepository "finpro-fenlie/src/user/user_repository"
	userUseCase "finpro-fenlie/src/user/user_use_case"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(v1Group *gin.RouterGroup, db *gorm.DB) {

	// Company Route
	companyRepository := companyRepository.NewCompanyRepository(db)
	companyUseCase := companyUseCase.NewCompanyUseCase(companyRepository)
	companyDelivery.NewCompanyDelivery(v1Group, companyUseCase)

	v1Group.Use(middleware.BasicAuth(companyUseCase))
	{
		// User Route
		userRepository := userRepository.NewUserRepository(db)
		userUseCase := userUseCase.NewUserUsecase(userRepository)
		userDelivery.NewUserDelivery(v1Group, userUseCase)

		// Category Route
		categoryRepo := categoryRepository.NewCategoryRepository(db)
		categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
		categoryDelivery.NewCategoryDelivery(v1Group, categoryUseCase)

		// Proudct Route
		productRepo := productRepository.NewProductRepository(db)
		productUseCase := productUseCase.NewProductUsecase(productRepo)
		productDelivery.NewProductDelivery(v1Group, productUseCase)

		// Transaction Route
		transactionRepo := transactionRepository.NewTransactionRepository(db)
		transactionUseCase := transactionUseCase.NewTransactionUsecase(transactionRepo)
		transactionDelivery.NewTransactionDelivery(v1Group, transactionUseCase)
	}

}
