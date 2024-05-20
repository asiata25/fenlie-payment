package user

import (
	"finpro-fenlie/model/dto/middlewareDto"
	"finpro-fenlie/model/dto/userDto"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	RetrieveLoginUser(*gin.Context, middlewareDto.LoginRequest, userDto.User) (string, error)
	InsertUser(*gin.Context, userDto.User, bool, bool) error
	RetrieveAllUser(*gin.Context, int, int, int64, string, string) (userDto.GetResponse, error)
	RetrieveUserByID(*gin.Context, string) (userDto.User, error)
	RetrieveUserByEmail(string) (userDto.User, error)
	CountUsers(*gin.Context, string, string) (int64, error)
	CheckUserEmailPassword(userDto.User) (bool, bool, error)
	EditUser(*gin.Context, string, map[string]interface{}) error
	RemoveUser(*gin.Context, string) error
}

type UserUsecase interface {
	Login(*gin.Context, middlewareDto.LoginRequest) (string, error)
	CreateUser(*gin.Context, userDto.User) error
	GetAllUser(*gin.Context, int, int, string, string) (userDto.GetResponse, error)
	GetUserByID(*gin.Context, string) (userDto.User, error)
	UpdateUser(*gin.Context, string, map[string]interface{}) error
	DeleteUser(*gin.Context, string) error
}
