package user

import (
	"finpro-fenlie/model/dto/user"
	"finpro-fenlie/model/entity"
)

type UserRepository interface {
	InsertUser(payload *entity.User) error
	// RetrieveAllUser(int, int, int64, string, string, *user.UserResponse) (user.UserResponse, error)
	// RetrieveUserByID(id string) (entity.User, error)
	RetrieveUserByEmail(string) (entity.User, error)
	// CountUsers(string, string, *user.UserResponse) (int64, error)
	// EditUser(payload *entity.User) error
	// RemoveUser(id string) error
}

type UserUseCase interface {
	Login(request user.LoginRequest) (string, error)
	CreateUser(request user.CreateUserRequest) error
	// GetAllUser(int, int, string, string) (user.UserResponse, error)
	// GetUserByID(string) (user.UserResponse, error)
	// UpdateUser(user.UpdateUserRequest) error
	// DeleteUser(string) error
}
