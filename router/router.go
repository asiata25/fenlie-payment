package router

import (
	"finpro-fenlie/src/category/categoryDelivery"
	"finpro-fenlie/src/category/categoryRepository"
	"finpro-fenlie/src/category/categoryUseCase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(v1Group *gin.RouterGroup, db *gorm.DB) {
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryUseCase := categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryDelivery.NewCategoryDelivery(v1Group, categoryUseCase)

}
