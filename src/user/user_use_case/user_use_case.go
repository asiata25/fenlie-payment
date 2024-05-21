package userUsecase

import (
	"finpro-fenlie/model/dto/auth"
	userDTO "finpro-fenlie/model/dto/user"
	"finpro-fenlie/src/user"

	"github.com/gin-gonic/gin"
)

type userUC struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserUsecase {
	return &userUC{userRepo}
}

// Implement Login
func (usecase *userUC) Login(c *gin.Context, req auth.LoginRequest) (string, error) {
	user, err := usecase.userRepo.RetrieveUserByEmail(req.Email)
	if err != nil {
		return "err", err
	}

	users, err := usecase.userRepo.RetrieveLoginUser(c, req, user)
	if err != nil {
		return "err", err
	}

	return users, nil
}

// Implement CreateUser
func (usecase *userUC) CreateUser(c *gin.Context, user userDTO.User) error {
	checkEmail, checkPass, err := usecase.userRepo.CheckUserEmailPassword(user)
	if err != nil {
		return err
	}

	err = usecase.userRepo.InsertUser(c, user, checkEmail, checkPass)
	if err != nil {
		return err
	}

	return nil
}

// Implement GetUser
func (usecase *userUC) GetAllUser(c *gin.Context, page, size int, email, name string) (userDTO.GetResponse, error) {
	totalData, err := usecase.userRepo.CountUsers(c, email, name)
	if err != nil {
		return userDTO.GetResponse{}, err
	}

	response, err := usecase.userRepo.RetrieveAllUser(c, page, size, totalData, email, name)
	if err != nil {
		return response, err
	}
	return response, nil
}

// Implement GetUserByID
func (usecase *userUC) GetUserByID(c *gin.Context, id string) (userDTO.User, error) {
	user, err := usecase.userRepo.RetrieveUserByID(c, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Implement UpdateUser
func (usecase *userUC) UpdateUser(c *gin.Context, id string, userUpdates map[string]interface{}) error {
	err := usecase.userRepo.EditUser(c, id, userUpdates)
	if err != nil {
		return err
	}
	return nil
}

// Implement DeleteUser
func (usecase *userUC) DeleteUser(c *gin.Context, id string) error {
	err := usecase.userRepo.RemoveUser(c, id)
	if err != nil {
		return err
	}
	return nil
}
