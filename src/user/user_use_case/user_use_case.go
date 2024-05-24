package userUseCase

import (
	"finpro-fenlie/helper"
	userDTO "finpro-fenlie/model/dto/user"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/user"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type userUC struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) user.UserUseCase {
	return &userUC{userRepo}
}

// Implement Login
func (usecase *userUC) Login(request userDTO.LoginRequest) (string, error) {
	user, err := usecase.userRepo.RetrieveUserByEmail(request.Email)
	if err != nil {
		return "err", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("invalid Password")
	}

	token, err := helper.GenerateTokenJwt(user.Email, user.Role, user.CompanyID, 10)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Implement CreateUser
func (usecase *userUC) CreateUser(request userDTO.CreateUserRequest) error {
	_, err := usecase.userRepo.RetrieveUserByEmail(request.Email)
	if err == nil {
		return errors.New("email is already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Role:     request.Role,
		Password: string(hash),
	}

	err = usecase.userRepo.InsertUser(&user)
	if err != nil {
		return err
	}

	return nil
}

// // Implement GetUser
// func (usecase *userUC) GetAllUser(c *gin.Context, page, size int, email, name string) (userDTO.GetResponse, error) {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return userDTO.GetResponse{}, err
// 	}

// 	totalData, err := usecase.userRepo.CountUsers(email, name, userInfo)
// 	if err != nil {
// 		return userDTO.GetResponse{}, err
// 	}

// 	response, err := usecase.userRepo.RetrieveAllUser(page, size, totalData, email, name, userInfo)
// 	if err != nil {
// 		return response, err
// 	}
// 	return response, nil
// }

// // Implement GetUserByID
// func (usecase *userUC) GetUserByID(c *gin.Context, id string) (userDTO.User, error) {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return userDTO.User{}, err
// 	}

// 	user, err := usecase.userRepo.RetrieveUserByID(id, userInfo)
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// // Implement UpdateUser
// func (usecase *userUC) UpdateUser(c *gin.Context, id string, userUpdates map[string]interface{}) error {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return err
// 	}

// 	err = usecase.userRepo.EditUser(id, userUpdates, userInfo)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Implement DeleteUser
// func (usecase *userUC) DeleteUser(c *gin.Context, id string) error {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return err
// 	}

// 	err = usecase.userRepo.RemoveUser(id, userInfo)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
