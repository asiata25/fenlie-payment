package companyUseCase

// import (
// 	"errors"
// 	companyDTO "finpro-fenlie/model/dto/company"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type companyMockUC struct {
// 	mock.Mock
// }

// func (uc *companyMockUC) GetAll() ([]*companyDTO.CompanyResponse, error) {
// 	args := uc.Called()
// 	return args.Get(0).([]*companyDTO.CompanyResponse), args.Error(1)
// }

// func (uc *companyMockUC) Create(request companyDTO.CompanyCreateRequest) error {
// 	args := uc.Called(request)
// 	return args.Error(0)
// }

// func (uc *companyMockUC) Delete(id string) error {
// 	args := uc.Called(id)
// 	return args.Error(0)
// }

// func (uc *companyMockUC) GetById(id string) (*companyDTO.CompanyResponse, error) {
// 	args := uc.Called(id)
// 	return args.Get(0).(*companyDTO.CompanyResponse), args.Error(1)
// }

// func (uc *companyMockUC) Update(request companyDTO.CompanyUpdateRequest) error {
// 	args := uc.Called(request)
// 	return args.Error(0)
// }

// func TestGetAll_Success(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)

// 	mockCompanyUC.On("GetAll").Return([]*companyDTO.CompanyResponse{}, nil)

// 	got, err := mockCompanyUC.GetAll()

// 	assert.Equal(t, []*companyDTO.CompanyResponse{}, got)
// 	assert.Nil(t, err)
// }

// func TestGetAll_Fail(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	expectedError := errors.New("Internal Server Error")

// 	mockCompanyUC.On("GetAll").Return([]*companyDTO.CompanyResponse{}, expectedError)

// 	got, err := mockCompanyUC.GetAll()

// 	assert.Equal(t, []*companyDTO.CompanyResponse{}, got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestCreate_Success(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	request := companyDTO.CompanyCreateRequest{
// 		Name:      "Company 1",
// 		Email:     "Company1@mail.com",
// 		SecretKey: "Secret key",
// 	}

// 	mockCompanyUC.On("Create", request).Return(nil)

// 	err := mockCompanyUC.Create(request)

// 	assert.Nil(t, err)
// }

// func TestCreate_Fail(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	request := companyDTO.CompanyCreateRequest{
// 		Name:      "Company 1",
// 		Email:     "Company1@mail.com",
// 		SecretKey: "Secret key",
// 	}
// 	expectedError := errors.New("Failed to create company")

// 	mockCompanyUC.On("Create", request).Return(expectedError)

// 	err := mockCompanyUC.Create(request)

// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestDelete_Success(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	id := uuid.New().String()

// 	mockCompanyUC.On("Delete", id).Return(nil)

// 	err := mockCompanyUC.Delete(id)

// 	assert.Nil(t, err)
// }

// func TestDelete_Fail(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	id := uuid.New().String()
// 	expectedError := errors.New("Internal server error")

// 	mockCompanyUC.On("Delete", id).Return(expectedError)

// 	err := mockCompanyUC.Delete(id)

// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestGetById_Success(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	id := uuid.New().String()

// 	mockCompanyUC.On("GetById", id).Return(&companyDTO.CompanyResponse{}, nil)

// 	got, err := mockCompanyUC.GetById(id)

// 	assert.Equal(t, &companyDTO.CompanyResponse{}, got)
// 	assert.Nil(t, err)
// }

// func TestGetById_Fail(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	id := uuid.New().String()
// 	expectedError := errors.New("Data Not Found")

// 	mockCompanyUC.On("GetById", id).Return(&companyDTO.CompanyResponse{}, expectedError)

// 	got, err := mockCompanyUC.GetById(id)

// 	assert.Equal(t, &companyDTO.CompanyResponse{}, got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestUpdate_Success(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	request := companyDTO.CompanyUpdateRequest{
// 		Name:      "Company 1",
// 		SecretKey: "Secret Key",
// 	}

// 	mockCompanyUC.On("Update", request).Return(nil)

// 	err := mockCompanyUC.Update(request)

// 	assert.Nil(t, err)
// }

// func TestUpdate_Fail(t *testing.T) {
// 	mockCompanyUC := new(companyMockUC)
// 	request := companyDTO.CompanyUpdateRequest{
// 		Name:      "Company 1",
// 		SecretKey: "Secret Key",
// 	}
// 	expectedError := errors.New("Failed to update company")

// 	mockCompanyUC.On("Update", request).Return(expectedError)

// 	err := mockCompanyUC.Update(request)

// 	assert.EqualError(t, err, expectedError.Error())
// }
