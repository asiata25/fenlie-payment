package companyRepository

import (
	"errors"
	"finpro-fenlie/model/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type companyRepoMock struct {
	mock.Mock
}

func (m *companyRepoMock) FindAll() ([]*entity.Company, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Company), args.Error(1)
}

func (m *companyRepoMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *companyRepoMock) RetrieveByID(id string) (*entity.Company, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Company), args.Error(1)
}

func (m *companyRepoMock) Save(payload entity.Company) error {
	args := m.Called(payload)
	return args.Error(0)
}

func (m *companyRepoMock) Update(payload entity.Company) error {
	args := m.Called(payload)
	return args.Error(0)
}

func TestFindAll_Success(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)

	mockRepoCompany.On("FindAll").Return([]*entity.Company{}, nil)

	got, err := mockRepoCompany.FindAll()

	assert.Equal(t, []*entity.Company{}, got)
	assert.Nil(t, err)
}

func TestFindAll_Fail(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	companies := []*entity.Company{}
	expectedError := errors.New("Internal Server Error")

	mockRepoCompany.On("FindAll").Return(companies, expectedError)

	got, err := mockRepoCompany.FindAll()

	assert.Equal(t, companies, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestDelete_Success(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	id := uuid.New().String()

	mockRepoCompany.On("Delete", id).Return(nil)

	err := mockRepoCompany.Delete(id)

	assert.Nil(t, err)
}

func TestDelete_Fail(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	id := uuid.New().String()
	expectedError := errors.New("Internal server error")

	mockRepoCompany.On("Delete", id).Return(expectedError)

	err := mockRepoCompany.Delete(id)

	assert.EqualError(t, err, expectedError.Error())
}

func TestRetrieveByID_Success(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	id := uuid.New().String()

	mockRepoCompany.On("RetrieveByID", id).Return(&entity.Company{}, nil)

	got, err := mockRepoCompany.RetrieveByID(id)

	assert.Equal(t, &entity.Company{}, got)
	assert.Nil(t, err)
}

func TestRetrieveByID_Fail(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	id := uuid.New().String()
	expectedError := errors.New("Data Not Found")

	mockRepoCompany.On("RetrieveByID", id).Return(&entity.Company{}, expectedError)

	got, err := mockRepoCompany.RetrieveByID(id)

	assert.Equal(t, &entity.Company{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestSave_Success(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	payload := entity.Company{
		Name:         "Company 1",
		Email:        "company1@mail.com",
		ClientSecret: "client secret",
	}

	mockRepoCompany.On("Save", payload).Return(nil)

	err := mockRepoCompany.Save(payload)

	assert.Nil(t, err)
}

func TestSave_Fail(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	payload := entity.Company{
		Name:         "Company 1",
		Email:        "company1@mail.com",
		ClientSecret: "client secret",
	}
	expectedError := errors.New("Failed to save company")

	mockRepoCompany.On("Save", payload).Return(expectedError)

	err := mockRepoCompany.Save(payload)

	assert.EqualError(t, err, expectedError.Error())
}

func TestUpdate_Success(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	payload := entity.Company{
		Name:         "Company 1",
		ClientSecret: "client secret",
	}

	mockRepoCompany.On("Update", payload).Return(nil)

	err := mockRepoCompany.Update(payload)

	assert.Nil(t, err)
}

func TestUpdate_Fail(t *testing.T) {
	mockRepoCompany := new(companyRepoMock)
	payload := entity.Company{
		Name:         "Company 1",
		ClientSecret: "client secret",
	}
	expectedError := errors.New("Failed to update company")

	mockRepoCompany.On("Update", payload).Return(expectedError)

	err := mockRepoCompany.Update(payload)

	assert.EqualError(t, err, expectedError.Error())
}
