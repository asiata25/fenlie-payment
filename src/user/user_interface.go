package user

import (
	userDTO "finpro-fenlie/model/dto/user"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	RetrieveLoginUser(*gin.Context, userDTO.LoginRequest, userDTO.User) (string, error)
	InsertUser(*gin.Context, userDTO.User, bool, bool) error
	RetrieveAllUser(*gin.Context, int, int, int64, string, string) (userDTO.GetResponse, error)
	RetrieveUserByID(*gin.Context, string) (userDTO.User, error)
	RetrieveUserByEmail(string) (userDTO.User, error)
	CountUsers(*gin.Context, string, string) (int64, error)
	CheckUserEmailPassword(userDTO.User) (bool, bool, error)
	EditUser(*gin.Context, string, map[string]interface{}) error
	RemoveUser(*gin.Context, string) error
}

type UserUsecase interface {
	Login(*gin.Context, userDTO.LoginRequest) (string, error)
	CreateUser(*gin.Context, userDTO.User) error
	GetAllUser(*gin.Context, int, int, string, string) (userDTO.GetResponse, error)
	GetUserByID(*gin.Context, string) (userDTO.User, error)
	UpdateUser(*gin.Context, string, map[string]interface{}) error
	DeleteUser(*gin.Context, string) error
}
