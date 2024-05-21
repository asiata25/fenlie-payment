package userUsecase

import (
	"errors"
	"finpro-fenlie/model/dto/middlewareDto"
	"finpro-fenlie/model/dto/userDto"
	"finpro-fenlie/pkg/middleware"
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
func (usecase *userUC) Login(c *gin.Context, req middlewareDto.LoginRequest) (string, error) {
	user, err := usecase.userRepo.RetrieveUserByEmail(req.Email)
	if err != nil {
		return "err", err
	}

	match := usecase.userRepo.ComparePassword(req.Password, user.Password)
	if !match {
		return "", errors.New("invalid password")
	}

	companyID := user.CompanyID.String()

	token, err := middleware.GenerateTokenJwt(req.Email, user.Role, companyID, 60)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Implement CreateUser
func (usecase *userUC) CreateUser(c *gin.Context, user userDto.User) error {
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
func (usecase *userUC) GetAllUser(c *gin.Context, page, size int, email, name string) (userDto.GetResponse, error) {
	totalData, err := usecase.userRepo.CountUsers(c, email, name)
	if err != nil {
		return userDto.GetResponse{}, err
	}

	response, err := usecase.userRepo.RetrieveAllUser(c, page, size, totalData, email, name)
	if err != nil {
		return response, err
	}
	return response, nil
}

// Implement GetUserByID
func (usecase *userUC) GetUserByID(c *gin.Context, id string) (userDto.User, error) {
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
