package productUsecase

import (
	"errors"
	productDTO "finpro-fenlie/model/dto/product"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type productMockUC struct {
	mock.Mock
}

func (m *productMockUC) GetAllProducts(page, pageSize int) ([]productDTO.ProductResponse, int, int64, int64, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]productDTO.ProductResponse), args.Get(1).(int), args.Get(2).(int64), args.Get(3).(int64), args.Error(4)
}

func (m *productMockUC) CreateProduct(request productDTO.ProductCreateRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *productMockUC) GetProduct(id string) (productDTO.ProductResponse, error) {
	args := m.Called(id)
	return args.Get(0).(productDTO.ProductResponse), args.Error(1)
}

func (m *productMockUC) UpdateProduct(id string, product productDTO.ProductUpdateRequest) error {
	args := m.Called(id, product)
	return args.Error(0)
}

func (m *productMockUC) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllProducts_Success(t *testing.T) {
	productMockUC := new(productMockUC)
	page, pageSize := 1, 10

	productMockUC.On("GetAllProducts", page, pageSize).Return([]productDTO.ProductResponse{}, 0, int64(0), int64(0), nil)

	got, _, _, _, err := productMockUC.GetAllProducts(page, pageSize)

	assert.Equal(t, []productDTO.ProductResponse{}, got)
	assert.Nil(t, err)
}

func TestGetAllProducts_Fail(t *testing.T) {
	productMockUC := new(productMockUC)
	page, pageSize := 1, 10
	expectedError := errors.New("Internal Server Error")

	productMockUC.On("GetAllProducts", page, pageSize).Return([]productDTO.ProductResponse{}, 0, int64(0), int64(0), expectedError)

	got, _, _, _, err := productMockUC.GetAllProducts(page, pageSize)

	assert.Equal(t, []productDTO.ProductResponse{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestCreateProduct_Success(t *testing.T) {
	productMockUC := new(productMockUC)
	request := productDTO.ProductCreateRequest{
		Name:        "Produk 1",
		Price:       5000,
		Description: "description",
		Status:      true,
		CategoryID:  "category_id",
		CompanyID:   "company_id",
	}

	productMockUC.On("CreateProduct", request).Return(nil)

	err := productMockUC.CreateProduct(request)

	assert.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
	productMockUC := new(productMockUC)
	request := productDTO.ProductCreateRequest{
		Name:        "Produk 1",
		Price:       5000,
		Description: "description",
		Status:      true,
		CategoryID:  "category_id",
		CompanyID:   "company_id",
	}
	expectedError := errors.New("Failed to create product")

	productMockUC.On("CreateProduct", request).Return(expectedError)

	err := productMockUC.CreateProduct(request)

	assert.EqualError(t, err, expectedError.Error())
}

func TestGetProduct_Success(t *testing.T) {
	productMockUC := new(productMockUC)
	id := uuid.New().String()

	productMockUC.On("GetProduct", id).Return(productDTO.ProductResponse{}, nil)

	got, err := productMockUC.GetProduct(id)

	assert.Equal(t, productDTO.ProductResponse{}, got)
	assert.Nil(t, err)
}

func TestGetProduct_Fail(t *testing.T) {
	productMockUC := new(productMockUC)
	id := uuid.New().String()
	expectedError := errors.New("Data not found")

	productMockUC.On("GetProduct", id).Return(productDTO.ProductResponse{}, expectedError)

	got, err := productMockUC.GetProduct(id)

	assert.Equal(t, productDTO.ProductResponse{}, got)
	assert.EqualError(t, err, expectedError.Error())
}

func TestUpdateProduct_Success(t *testing.T) {
	productRepoMock := new(productMockUC)
	id := uuid.New().String()
	product := productDTO.ProductUpdateRequest{
		Name:        "Produk 1",
		Price:       5000,
		Description: "description",
		Status:      true,
		CategoryID:  "category_id",
		CompanyID:   "company_id",
	}

	productRepoMock.On("UpdateProduct", id, product).Return(nil)

	err := productRepoMock.UpdateProduct(id, product)

	assert.Nil(t, err)
}

func TestUpdateProduct_Fail(t *testing.T) {
	productRepoMock := new(productMockUC)
	id := uuid.New().String()
	product := productDTO.ProductUpdateRequest{
		Name:        "Produk 1",
		Price:       5000,
		Description: "description",
		Status:      true,
		CategoryID:  "category_id",
		CompanyID:   "company_id",
	}
	expectedError := errors.New("Failed to update product")

	productRepoMock.On("UpdateProduct", id, product).Return(expectedError)

	err := productRepoMock.UpdateProduct(id, product)

	assert.EqualError(t, err, expectedError.Error())
}

func TestDeleteProduct_Success(t *testing.T) {
	productMockUC := new(productMockUC)
	id := uuid.New().String()

	productMockUC.On("DeleteProduct", id).Return(nil)

	err := productMockUC.DeleteProduct(id)

	assert.Nil(t, err)
}

func TestDeleteProduct_Fail(t *testing.T) {
	productMockUC := new(productMockUC)
	id := uuid.New().String()
	expectedError := errors.New("Internal Server Error")

	productMockUC.On("DeleteProduct", id).Return(expectedError)

	err := productMockUC.DeleteProduct(id)

	assert.EqualError(t, err, expectedError.Error())
}
