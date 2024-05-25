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
	user, err := usecase.userRepo.RetrieveUserByEmail(request.Email, request.CompanyID)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := helper.GenerateTokenJwt(user.ID, user.Role, 10)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Implement CreateUser
func (usecase *userUC) CreateUser(request userDTO.CreateUserRequest) error {
	_, err := usecase.userRepo.RetrieveUserByEmail(request.Email, request.CompanyID)
	if err == nil {
		return errors.New("email is already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Name:      request.Name,
		Email:     request.Email,
		Role:      request.Role,
		Password:  string(hash),
		CompanyID: request.CompanyID,
	}

	err = usecase.userRepo.InsertUser(&user)
	if err != nil {
		return err
	}

	return nil
}

// Implement GetUser
func (usecase *userUC) GetAllUser(page, size int, email, name, companyId string) ([]userDTO.UserResponse, int, error) {
	var users []userDTO.UserResponse
	response, total, err := usecase.userRepo.RetrieveAllUser(page, size, email, name, companyId)

	for i, resp := range response {
		users = append(users, helper.ToUserResponse(resp))
		users[i].Company = resp.Company.Name
	}

	if err != nil {
		return users, total, err
	}

	return users, total, nil
}

// Implement GetUserByID
func (usecase *userUC) GetUserByID(id, companyId string) (userDTO.UserResponse, error) {

	result, err := usecase.userRepo.RetrieveUserByID(id, companyId)
	if err != nil {
		return userDTO.UserResponse{}, err
	}

	user := helper.ToUserResponse(result)
	user.Company = result.Company.Name

	return user, nil
}

// Implement UpdateUser
func (usecase *userUC) UpdateUser(request userDTO.UpdateUserRequest) error {
	user := entity.User{
		ID:        request.ID,
		Name:      request.Name,
		Role:      request.Role,
		CompanyID: request.CompanyID,
		Password:  request.Password,
	}

	userExisting, err := usecase.userRepo.RetrieveUserByID(user.ID, user.CompanyID)
	if err != nil {
		return err
	}

	if request.Password == "" {
		request.Password = userExisting.Password
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	err = usecase.userRepo.EditUser(&user)
	if err != nil {
		return err
	}
	return nil
}

// Implement DeleteUser
func (usecase *userUC) DeleteUser(id, companyId string) error {
	err := usecase.userRepo.RemoveUser(id, companyId)
	return err
}
