package userUsecase

import (
	"errors"
	"finpro-fenlie/model/dto/auth"
	userDTO "finpro-fenlie/model/dto/user"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userUCMock struct {
	mock.Mock
}

func (m *userUCMock) Login(c *gin.Context, req auth.LoginRequest) (string, error) {
	args := m.Called(c, req)
	return args.String(0), args.Error(1)
}

func (m *userUCMock) CreateUser(c *gin.Context, user userDTO.User) error {
	args := m.Called(c, user)
	return args.Error(0)
}

func (m *userUCMock) GetAllUser(c *gin.Context, page, size int, email, name string) (userDTO.GetResponse, error) {
	args := m.Called(c, page, size, email, name)
	return args.Get(0).(userDTO.GetResponse), args.Error(1)
}

func (m *userUCMock) GetUserByID(c *gin.Context, id string) (userDTO.User, error) {
	args := m.Called(c, id)
	return args.Get(0).(userDTO.User), args.Error(1)
}

func (m *userUCMock) UpdateUser(c *gin.Context, id string, data map[string]interface{}) error {
	args := m.Called(c, id, data)
	return args.Error(0)
}

func (m *userUCMock) DeleteUser(c *gin.Context, id string) error {
	args := m.Called(c, id)
	return args.Error(0)
}

// Start of test
func TestLogin_Success(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	req := auth.LoginRequest{
		Email:    "user@gmail.com",
		Password: "password",
	}
	mockUserUC.On("Login", c, req).Return("token", nil)

	result, err := mockUserUC.Login(c, req)

	assert.Nil(t, err)
	assert.Equal(t, "token", result)
}

func TestLogin_Failed(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	req := auth.LoginRequest{
		Email:    "user@mail.com",
		Password: "password",
	}
	expectedError := errors.New("invalid email or password")

	mockUserUC.On("Login", c, req).Return("", expectedError)

	result, err := mockUserUC.Login(c, req)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, "", result)
}

func TestCreateUser_Success(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	user := userDTO.User{
		ID:        uuid.New(),
		Name:      "User Name",
		Email:     "user@gmail.com",
		Password:  "password",
		CompanyID: uuid.New(),
		Role:      "USER",
	}

	mockUserUC.On("CreateUser", c, user).Return(nil)

	err := mockUserUC.CreateUser(c, user)

	assert.Nil(t, err)
}

func TestCreateUser_Failed(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	user := userDTO.User{
		ID:        uuid.New(),
		Name:      "User Name",
		Email:     "user@gmail.com",
		Password:  "password",
		CompanyID: uuid.New(),
		Role:      "USER",
	}
	expectedError := errors.New("email already exist")

	mockUserUC.On("CreateUser", c, user).Return(expectedError)

	err := mockUserUC.CreateUser(c, user)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}

func TestGetAllUser_Success(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	page := 1
	size := 10
	totalData := 5
	email, name := "user@gmail.com", "User Name"
	data := userDTO.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  "password",
		CompanyID: uuid.New(),
		Role:      "USER",
	}
	pagination := userDTO.Paging{
		Page: page,
		Size: size,
	}
	response := userDTO.GetResponse{
		Data:       []userDTO.User{data},
		Pagination: pagination,
		TotalData:  int64(totalData),
	}

	mockUserUC.On("GetAllUser", c, page, size, email, name).Return(response, nil)

	result, err := mockUserUC.GetAllUser(c, page, size, email, name)

	assert.Nil(t, err)
	assert.Equal(t, response, result)
}

func TestGetAllUser_Failed(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	page := 1
	size := 10
	email, name := "user@gmail.com", "User Name"
	data := userDTO.User{}
	pagination := userDTO.Paging{
		Page: page,
		Size: size,
	}
	response := userDTO.GetResponse{
		Data:       []userDTO.User{data},
		Pagination: pagination,
		TotalData:  0,
	}
	expectedError := errors.New("failed to retrieve data")

	mockUserUC.On("GetAllUser", c, page, size, email, name).Return(response, expectedError)

	result, err := mockUserUC.GetAllUser(c, page, size, email, name)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, response, result)
}

func TestGetUserByID_Success(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	id := uuid.New()
	data := userDTO.User{
		ID:        uuid.New(),
		Name:      "User Name",
		Email:     "user@gmail.com",
		Password:  "password",
		CompanyID: uuid.New(),
		Role:      "USER",
	}

	mockUserUC.On("GetUserByID", c, id.String()).Return(data, nil)

	result, err := mockUserUC.GetUserByID(c, id.String())

	assert.Nil(t, err)
	assert.Equal(t, data, result)
}

func TestGetUserByID_Failed(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	id := uuid.New()
	expectedError := errors.New("failed to retrieve data")

	mockUserUC.On("GetUserByID", c, id.String()).Return(userDTO.User{}, expectedError)

	result, err := mockUserUC.GetUserByID(c, id.String())

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, userDTO.User{}, result)
}

func TestUpdateUser_Success(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	id := uuid.New()
	data := map[string]interface{}{
		"name": "User Name",
	}

	mockUserUC.On("UpdateUser", c, id.String(), data).Return(nil)

	err := mockUserUC.UpdateUser(c, id.String(), data)

	assert.Nil(t, err)
}

func TestUpdateUser_Failed(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	id := uuid.New()
	data := map[string]interface{}{
		"name": "User Name",
	}
	expectedError := errors.New("failed to update data")

	mockUserUC.On("UpdateUser", c, id.String(), data).Return(expectedError)

	err := mockUserUC.UpdateUser(c, id.String(), data)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}

func TestDeleteUser_Success(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	id := uuid.New()

	mockUserUC.On("DeleteUser", c, id.String()).Return(nil)

	err := mockUserUC.DeleteUser(c, id.String())

	assert.Nil(t, err)
}

func TestDeleteUser_Failed(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	mockUserUC := new(userUCMock)
	id := uuid.New()
	expectedError := errors.New("failed to delete data")

	mockUserUC.On("DeleteUser", c, id.String()).Return(expectedError)

	err := mockUserUC.DeleteUser(c, id.String())

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}
