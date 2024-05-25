package user

import (
	"finpro-fenlie/model/dto/user"
	"finpro-fenlie/model/entity"
)

type UserRepository interface {
	InsertUser(payload *entity.User) error
	RetrieveAllUser(page, size int, email, name, companyId string) ([]entity.User, int, error)
	RetrieveUserByID(id, companyId string) (entity.User, error)
	RetrieveUserByEmail(id, companyId string) (entity.User, error)
	// CountUsers(string, string, *user.UserResponse) (int64, error)
	EditUser(payload *entity.User) error
	RemoveUser(id, companyId string) error
}

type UserUseCase interface {
	Login(request user.LoginRequest) (string, error)
	CreateUser(request user.CreateUserRequest) error
	GetAllUser(page, size int, email, name, companyId string) ([]user.UserResponse, int, error)
	GetUserByID(id, companyId string) (user.UserResponse, error)
	UpdateUser(request user.UpdateUserRequest) error
	DeleteUser(id, companyId string) error
}
