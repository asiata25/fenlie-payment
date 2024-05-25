package userRepository

// import (
// 	"errors"
// 	"finpro-fenlie/model/dto/middlewareDto"
// 	"finpro-fenlie/model/dto/userDto"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type userRepoMock struct {
// 	mock.Mock
// }

// func (m *userRepoMock) RetrieveLoginUser(ctx *gin.Context, req middlewareDto.LoginRequest, user userDto.User) (string, error) {
// 	args := m.Called(ctx, req, user)
// 	return args.String(0), args.Error(1)
// }

// func (m *userRepoMock) InsertUser(c *gin.Context, user userDto.User, checkEmail, checkPass bool) error {
// 	args := m.Called(c, user, checkEmail, checkPass)
// 	return args.Error(0)
// }

// func (m *userRepoMock) RetrieveAllUser(c *gin.Context, page, size int, totalData int64, email, name string) (userDto.GetResponse, error) {
// 	args := m.Called(c, page, size, totalData, email, name)
// 	return args.Get(0).(userDto.GetResponse), args.Error(1)
// }

// func (m *userRepoMock) RetrieveUserByID(c *gin.Context, id string) (userDto.User, error) {
// 	args := m.Called(c, id)
// 	return args.Get(0).(userDto.User), args.Error(1)
// }

// func (m *userRepoMock) EditUser(c *gin.Context, id string, userUpdates map[string]interface{}) error {
// 	args := m.Called(c, id, userUpdates)
// 	return args.Error(0)
// }

// func (m *userRepoMock) RemoveUser(c *gin.Context, id string) error {
// 	args := m.Called(c, id)
// 	return args.Error(0)
// }

// func (m *userRepoMock) RetrieveUserByEmail(email string) (userDto.User, error) {
// 	args := m.Called(email)
// 	return args.Get(0).(userDto.User), args.Error(1)
// }

// func (m *userRepoMock) CountUsers(c *gin.Context, email, name string) (int64, error) {
// 	args := m.Called(c, email, name)
// 	return args.Get(0).(int64), args.Error(1)
// }

// func (m *userRepoMock) CheckUserEmailPassword(user userDto.User) (bool, bool, error) {
// 	args := m.Called(user)
// 	return args.Bool(0), args.Bool(1), args.Error(2)
// }

// // Start Testing
// func TestRetrieveLoginUser_Success(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	req := middlewareDto.LoginRequest{Email: "user@gmail.com", Password: "password"}
// 	user := userDto.User{Email: "user@gmail.com", Password: "password"}
// 	userRepoMock.On("RetrieveLoginUser", c, req, user).Return("token", nil)
// 	got, err := userRepoMock.RetrieveLoginUser(c, req, user)

// 	assert.Nil(t, err)
// 	assert.Equal(t, "token", got)
// }

// func TestRetrieveLoginUser_Fail(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	req := middlewareDto.LoginRequest{Email: "usar@gmail.com", Password: "password"}
// 	user := userDto.User{Email: "user@gmail.com", Password: "password"}
// 	expectedError := errors.New("invalid password")

// 	userRepoMock.On("RetrieveLoginUser", c, req, user).Return("", expectedError)

// 	got, err := userRepoMock.RetrieveLoginUser(c, req, user)

// 	assert.Equal(t, "", got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestInsertUser_Success(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	id := uuid.New()
// 	company := uuid.New()
// 	user := userDto.User{ID: id, Name: "User Name", Email: "user@gmail.com", Password: "password", CompanyID: company}
// 	checkEmail, checkPass := true, true

// 	userRepoMock.On("InsertUser", c, user, checkEmail, checkPass).Return(nil)

// 	err := userRepoMock.InsertUser(c, user, checkEmail, checkPass)

// 	assert.Nil(t, err)
// }

// func TestInsertUser_Fail(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	user := userDto.User{Email: "user@gmail.com", Password: "password"}
// 	checkEmail, checkPass := true, true

// 	expectedError := errors.New("email or password validation failed")

// 	userRepoMock.On("InsertUser", c, user, checkEmail, checkPass).Return()

// 	err := userRepoMock.InsertUser(c, user, checkEmail, checkPass)

// 	assert.EqualError(t, err, expectedError.Error())
// 	userRepoMock.AssertExpectations(t)
// }

// func TestRetrieveAllUser_Success(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	page, size := 1, 10
// 	totalData := int64(1)
// 	email, name := "user@gmail.com", "User Name"
// 	id := uuid.New()
// 	company := uuid.New()
// 	user := userDto.User{ID: id, Name: name, Email: email, Password: "password", CompanyID: company}
// 	response := userDto.GetResponse{Data: []userDto.User{user}, Pagination: userDto.Paging{Page: page, Size: size}, TotalData: totalData}

// 	userRepoMock.On("RetrieveAllUser", c, page, size, totalData, email, name).Return(response, nil)

// 	got, err := userRepoMock.RetrieveAllUser(c, page, size, totalData, email, name)

// 	assert.Nil(t, err)
// 	assert.Equal(t, response, got)
// }

// func TestRetrieveAllUser_Fail(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	page, size := 1, 10
// 	totalData := int64(1)
// 	email, name := "user@gmail.com", "User Name"
// 	expectedError := errors.New("internal server error")

// 	userRepoMock.On("RetrieveAllUser", c, page, size, totalData, email, name).Return(userDto.GetResponse{}, expectedError)

// 	got, err := userRepoMock.RetrieveAllUser(c, page, size, totalData, email, name)

// 	assert.Equal(t, userDto.GetResponse{}, got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestRetrieveUserByID_Success(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	id := uuid.New()
// 	company := uuid.New()
// 	user := userDto.User{ID: id, Name: "User Name", Email: "user@gmail.com", Password: "password", CompanyID: company}

// 	userRepoMock.On("RetrieveUserByID", c, id.String()).Return(user, nil)

// 	got, err := userRepoMock.RetrieveUserByID(c, id.String())

// 	assert.Nil(t, err)
// 	assert.Equal(t, user, got)
// }

// func TestRetrieveUserByID_Fail(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	id := uuid.New()
// 	expectedError := errors.New("internal server error")

// 	userRepoMock.On("RetrieveUserByID", c, id.String()).Return(userDto.User{}, expectedError)

// 	got, err := userRepoMock.RetrieveUserByID(c, id.String())

// 	assert.Equal(t, userDto.User{}, got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestEditUser_Success(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	userUpdates := map[string]interface{}{"name": "User Name", "password": "password"}
// 	user := userDto.User{ID: uuid.New(), Name: "User Name", Email: "user@gmail.com", Password: "password", CompanyID: uuid.New(), Role: "ADMIN"}

// 	userRepoMock.On("EditUser", c, user.ID.String(), userUpdates).Return(nil)

// 	err := userRepoMock.EditUser(c, user.ID.String(), userUpdates)

// 	assert.Nil(t, err)
// }

// func TestEditUser_Fail(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	userUpdates := map[string]interface{}{"name": "User Name", "password": "password"}
// 	user := userDto.User{ID: uuid.New(), Name: "User Name", Email: "user@gmail.com", Password: "password", CompanyID: uuid.New(), Role: "ADMIN"}
// 	expectedError := errors.New("internal server error")

// 	userRepoMock.On("EditUser", c, user.ID.String(), userUpdates).Return(expectedError)

// 	err := userRepoMock.EditUser(c, user.ID.String(), userUpdates)

// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestRemoveUser_Success(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	id := uuid.New()

// 	userRepoMock.On("RemoveUser", c, id.String()).Return(nil)

// 	err := userRepoMock.RemoveUser(c, id.String())

// 	assert.Nil(t, err)
// }

// func TestRemoveUser_Fail(t *testing.T) {
// 	c, _ := gin.CreateTestContext(nil)
// 	userRepoMock := new(userRepoMock)
// 	id := uuid.New()
// 	expectedError := errors.New("unauthorized")

// 	userRepoMock.On("RemoveUser", c, id.String()).Return(expectedError)

// 	err := userRepoMock.RemoveUser(c, id.String())

// 	assert.EqualError(t, err, expectedError.Error())
// }
