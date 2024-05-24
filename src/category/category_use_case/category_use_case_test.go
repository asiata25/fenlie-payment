package categoryUseCase

import (
	"errors"
	categoryDto "finpro-fenlie/model/dto/category"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type categoryMock struct {
	mock.Mock
}

func (m *categoryMock) CreateLoan(request *categoryDto.CategoryRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *categoryMock) GetAllLoans(page, size string) (*[]categoryDto.CategoryResponse, error) {
	args := m.Called(page, size)
	return args.Get(0).(*[]categoryDto.CategoryResponse), args.Error(1)
}

func (m *categoryMock) GetLoanById(ID string) (categoryDto.CategoryResponse, error) {
	args := m.Called(ID)
	return args.Get(0).(categoryDto.CategoryResponse), args.Error(1)
}

func (m *categoryMock) UpdateLoan(request *categoryDto.CategoryRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *categoryMock) DeleteLoan(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func TestCreateLoan_Success(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	category := categoryDto.CategoryRequest{
		Name:      "Category 1",
		CompanyID: uuid.New().String(),
	}

	mockCategoryUC.On("CreateLoan", &category).Return(nil)

	err := mockCategoryUC.CreateLoan(&category)

	assert.Nil(t, err)
}

func TestCreateLoan_Fail(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	category := categoryDto.CategoryRequest{
		Name:      "Category 1",
		CompanyID: uuid.New().String(),
	}

	expectedError := errors.New("Failed to create category")

	mockCategoryUC.On("CreateLoan", &category).Return(expectedError)

	err := mockCategoryUC.CreateLoan(&category)

	assert.EqualError(t, err, expectedError.Error())
}

func TestGetAllLoans_Success(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	page, size := "1", "10"

	mockCategoryUC.On("GetAllLoans", page, size).Return(&[]categoryDto.CategoryResponse{}, nil)

	got, err := mockCategoryUC.GetAllLoans(page, size)

	assert.Equal(t, &[]categoryDto.CategoryResponse{}, got)
	assert.Nil(t, err)
}

func TestGetAllLoans_Fail(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	page, size := "1", "10"

	expectedError := errors.New("Internal Server Error")

	mockCategoryUC.On("GetAllLoans", page, size).Return(&[]categoryDto.CategoryResponse{}, expectedError)

	got, err := mockCategoryUC.GetAllLoans(page, size)

	assert.Equal(t, &[]categoryDto.CategoryResponse{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetLoaById_Success(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	ID := uuid.New().String()

	mockCategoryUC.On("GetLoanById", ID).Return(categoryDto.CategoryResponse{}, nil)

	got, err := mockCategoryUC.GetLoanById(ID)

	assert.Equal(t, categoryDto.CategoryResponse{}, got)
	assert.Nil(t, err)
}

func TestGetLoaById_Fail(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	ID := uuid.New().String()
	expectedError := errors.New("Internal Server Error")

	mockCategoryUC.On("GetLoanById", ID).Return(categoryDto.CategoryResponse{}, expectedError)

	got, err := mockCategoryUC.GetLoanById(ID)

	assert.Equal(t, categoryDto.CategoryResponse{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestUpdateLoan_Success(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	category := categoryDto.CategoryRequest{
		Name:      "Category 1",
		CompanyID: uuid.New().String(),
	}

	mockCategoryUC.On("UpdateLoan", &category).Return(nil)

	err := mockCategoryUC.UpdateLoan(&category)

	assert.Nil(t, err)
}

func TestUpdateLoan_Fail(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	category := categoryDto.CategoryRequest{
		Name:      "Category 1",
		CompanyID: uuid.New().String(),
	}

	expectedError := errors.New("Failed update category")

	mockCategoryUC.On("UpdateLoan", &category).Return(expectedError)

	err := mockCategoryUC.UpdateLoan(&category)

	assert.EqualError(t, err, expectedError.Error())
}

func TestDeleteLoan_Success(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	ID := uuid.New().String()

	mockCategoryUC.On("DeleteLoan", ID).Return(nil)

	err := mockCategoryUC.DeleteLoan(ID)

	assert.Nil(t, err)
}

func TestDeleteLoan_Fail(t *testing.T) {
	mockCategoryUC := new(categoryMock)
	ID := uuid.New().String()

	expectedError := errors.New("Internal Server Error")

	mockCategoryUC.On("DeleteLoan", ID).Return(expectedError)

	err := mockCategoryUC.DeleteLoan(ID)

	assert.EqualError(t, err, expectedError.Error())
}
