package productRepository

// import (
// 	"errors"
// 	"finpro-fenlie/model/entity"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type productRepoMock struct {
// 	mock.Mock
// }

// func (m *productRepoMock) GetAllProducts(page, pageSize int) ([]entity.Product, error) {
// 	args := m.Called(page, pageSize)
// 	return args.Get(0).([]entity.Product), args.Error(1)
// }

// func (m *productRepoMock) InsertProduct(product entity.Product) error {
// 	args := m.Called(product)
// 	return args.Error(0)
// }

// func (m *productRepoMock) GetById(id string) (entity.Product, error) {
// 	args := m.Called(id)
// 	return args.Get(0).(entity.Product), args.Error(1)
// }

// func (m *productRepoMock) UpdateProduct(id string, product entity.Product) error {
// 	args := m.Called(id, product)
// 	return args.Error(0)
// }

// func (m *productRepoMock) SoftDeleteProduct(id string) error {
// 	args := m.Called(id)
// 	return args.Error(0)
// }

// func TestGetAllProducts_Success(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	page, pageSize := 1, 10

// 	productRepoMock.On("GetAllProducts", page, pageSize).Return([]entity.Product{}, nil)

// 	got, err := productRepoMock.GetAllProducts(page, pageSize)

// 	assert.Equal(t, []entity.Product{}, got)
// 	assert.Nil(t, err)
// }

// func TestGetAllProducts_Fail(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	page, pageSize := 1, 10
// 	expectedError := errors.New("Internal server error")

// 	productRepoMock.On("GetAllProducts", page, pageSize).Return([]entity.Product{}, expectedError)

// 	got, err := productRepoMock.GetAllProducts(page, pageSize)

// 	assert.Equal(t, []entity.Product{}, got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestInsertProduct_Success(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	product := entity.Product{
// 		Name:        "Produk 1",
// 		Price:       5000,
// 		Description: "description",
// 		Status:      true,
// 		CategoryID:  "category_id",
// 		CompanyID:   "company_id",
// 	}

// 	productRepoMock.On("InsertProduct", product).Return(nil)

// 	err := productRepoMock.InsertProduct(product)

// 	assert.Nil(t, err)
// }

// func TestInsertProduct_Fail(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	product := entity.Product{
// 		Name:        "Produk 1",
// 		Price:       5000,
// 		Description: "description",
// 		Status:      true,
// 		CategoryID:  "category_id",
// 		CompanyID:   "company_id",
// 	}
// 	expectedError := errors.New("Failed to create Product")

// 	productRepoMock.On("InsertProduct", product).Return(expectedError)

// 	err := productRepoMock.InsertProduct(product)

// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestGetById_Success(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	id := uuid.New().String()

// 	productRepoMock.On("GetById", id).Return(entity.Product{}, nil)

// 	got, err := productRepoMock.GetById(id)

// 	assert.Equal(t, entity.Product{}, got)
// 	assert.Nil(t, err)
// }

// func TestGetById_Fail(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	id := uuid.New().String()
// 	expectedError := errors.New("Data Not Found")

// 	productRepoMock.On("GetById", id).Return(entity.Product{}, expectedError)

// 	got, err := productRepoMock.GetById(id)

// 	assert.Equal(t, entity.Product{}, got)
// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestUpdateProduct_Success(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	id := uuid.New().String()
// 	product := entity.Product{
// 		Name:        "Produk 1",
// 		Price:       5000,
// 		Description: "description",
// 		Status:      true,
// 		CategoryID:  "category_id",
// 		CompanyID:   "company_id",
// 	}

// 	productRepoMock.On("UpdateProduct", id, product).Return(nil)

// 	err := productRepoMock.UpdateProduct(id, product)

// 	assert.Nil(t, err)
// }

// func TestUpdateProduct_Fail(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	id := uuid.New().String()
// 	product := entity.Product{
// 		Name:        "Produk 1",
// 		Price:       5000,
// 		Description: "description",
// 		Status:      true,
// 		CategoryID:  "category_id",
// 		CompanyID:   "company_id",
// 	}
// 	expectedError := errors.New("Failed to update product")

// 	productRepoMock.On("UpdateProduct", id, product).Return(expectedError)

// 	err := productRepoMock.UpdateProduct(id, product)

// 	assert.EqualError(t, err, expectedError.Error())
// }

// func TestSoftDeleteProduct_Success(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	id := uuid.New().String()

// 	productRepoMock.On("SoftDeleteProduct", id).Return(nil)

// 	err := productRepoMock.SoftDeleteProduct(id)

// 	assert.Nil(t, err)
// }

// func TestSoftDeleteProduct_Fail(t *testing.T) {
// 	productRepoMock := new(productRepoMock)
// 	id := uuid.New().String()
// 	expectedError := errors.New("Internal Server Error")

// 	productRepoMock.On("SoftDeleteProduct", id).Return(expectedError)

// 	err := productRepoMock.SoftDeleteProduct(id)

// 	assert.EqualError(t, err, expectedError.Error())
// }
