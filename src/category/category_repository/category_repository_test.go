package categoryRepository

import (
	"errors"
	"finpro-fenlie/model/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type categoryRepoMock struct {
	mock.Mock
}

func (m *categoryRepoMock) Save(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *categoryRepoMock) GetAll(page, size int) (*[]entity.Category, int64, error) {
	args := m.Called(page, size)
	return args.Get(0).(*[]entity.Category), args.Get(1).(int64), args.Error(2)
}

func (m *categoryRepoMock) GetById(id string) (entity.Category, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Category), args.Error(1)
}

func (m *categoryRepoMock) Update(category *entity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *categoryRepoMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestSave_Success(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	id := uuid.New().String()
	company := uuid.New().String()
	category := entity.Category{ID: id, Name: "Category 1", CompanyID: company}

	categoryRepoMock.On("Save", &category).Return(nil)

	err := categoryRepoMock.Save(&category)

	assert.Nil(t, err)
}

func TestSave_Fail(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	id := uuid.New().String()
	company := uuid.New().String()
	category := entity.Category{ID: id, Name: "Category 1", CompanyID: company}
	expectedError := errors.New("failed to save category")

	categoryRepoMock.On("Save", &category).Return(expectedError)

	err := categoryRepoMock.Save(&category)

	assert.EqualError(t, err, expectedError.Error())
}

func TestGetAll_Success(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	page, size := 1, 10

	categoryRepoMock.On("GetAll", page, size).Return(&[]entity.Category{}, int64(0), nil)

	got, _, err := categoryRepoMock.GetAll(page, size)

	assert.Equal(t, &[]entity.Category{}, got)
	assert.Nil(t, err)
}

func TestGetAll_Fail(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	page, size := 1, 10
	expectedError := errors.New("internal server error")

	categoryRepoMock.On("GetAll", page, size).Return(&[]entity.Category{}, int64(0), expectedError)

	got, _, err := categoryRepoMock.GetAll(page, size)

	assert.Equal(t, &[]entity.Category{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetById_Success(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	id := uuid.New().String()
	company := uuid.New().String()
	category := entity.Category{ID: id, Name: "category 1", CompanyID: company}

	categoryRepoMock.On("GetById", id).Return(category, nil)

	got, err := categoryRepoMock.GetById(id)

	assert.Nil(t, err)
	assert.Equal(t, category, got)
}

func TestGetById_Fail(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	id := uuid.New().String()
	expectedError := errors.New("internal server error")

	categoryRepoMock.On("GetById", id).Return(entity.Category{}, expectedError)

	got, err := categoryRepoMock.GetById(id)

	assert.Equal(t, entity.Category{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestUpdate_Success(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	categoryUpdate := entity.Category{Name: "Category Updated", CompanyID: uuid.New().String()}

	categoryRepoMock.On("Update", &categoryUpdate).Return(nil)

	err := categoryRepoMock.Update(&categoryUpdate)

	assert.Nil(t, err)
}

func TestUpdate_Fail(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	categoryUpdate := entity.Category{Name: "Category Updated", CompanyID: uuid.New().String()}
	expectedError := errors.New("failed to update category")

	categoryRepoMock.On("Update", &categoryUpdate).Return(expectedError)

	err := categoryRepoMock.Update(&categoryUpdate)

	assert.EqualError(t, err, expectedError.Error())
}

func TestDelete_Success(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	id := uuid.New().String()

	categoryRepoMock.On("Delete", id).Return(nil)

	err := categoryRepoMock.Delete(id)

	assert.Nil(t, err)
}

func TestDelete_Fail(t *testing.T) {
	categoryRepoMock := new(categoryRepoMock)
	id := uuid.New().String()
	expectedError := errors.New("Internal server error")

	categoryRepoMock.On("Delete", id).Return(expectedError)

	err := categoryRepoMock.Delete(id)

	assert.EqualError(t, err, expectedError.Error())
}
