package router

import (
	"finpro-fenlie/src/user/userDelivery"
	"finpro-fenlie/src/user/userRepository"
	"finpro-fenlie/src/user/userUsecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(v1Group *gin.RouterGroup, db *gorm.DB) {
	userRepo := userRepository.NewUserRepository(db)
	userUC := userUsecase.NewUserUsecase(userRepo)
	userDelivery.NewUserDelivery(v1Group, userUC)
}
