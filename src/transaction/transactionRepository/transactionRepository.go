package transactionRepository

import (
	"errors"
	"finpro-fenlie/model/dto/middlewareDto"
	"finpro-fenlie/model/dto/productDto"
	"finpro-fenlie/model/dto/transactionDto"
	"finpro-fenlie/model/dto/userDto"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/src/product/product_repository"
	"finpro-fenlie/src/transaction"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) InputTransaction(request transactionDto.RequestTransaction, userInfo *middlewareDto.UserInfo) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	idTrx := uuid.New()
	orderDate := time.Now().Format("2006-01-02 15:04:05")
	totalTrx := 0

	// Create Transaction
	err := r.db.Create(&entity.Transaction{
		ID:        idTrx,
		UserID:    uuid.MustParse(userInfo.ID),
		OrderDate: orderDate,
		Customer:  request.Customer,
		Status:    false,
		Total:     totalTrx,
		CompanyID: uuid.MustParse(userInfo.CompanyID),
	}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("failed to create transaction")
	}

	// Get Product Data, Create Detail Transaction and Invoices
	productRepo := product_repository.NewProductRepository(r.db)
	for _, v := range request.RequestProductEqual {
		product, err := productRepo.RetrieveByID(v.ProductID.String(), userInfo)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return errors.New("product with ID " + v.ProductID.String() + " not found")
			}
		}

		if !product.Status {
			tx.Rollback()
			return errors.New("product with ID " + v.ProductID.String() + " is not available")
		}

		total := product.Price * v.Quantity

		// Create Detail Transaction
		err = r.db.Create(&entity.DetailTransaction{
			ID:        uuid.New(),
			OrderID:   idTrx,
			ProductID: product.ID,
			Quantity:  v.Quantity,
			Total:     total,
			CompanyID: product.CompanyID,
		}).Error
		if err != nil {
			tx.Rollback()
			return errors.New("failed to create detail transaction")
		}
		totalTrx += total
	}

	// Create Invoices
	err = r.db.Create(&entity.Invoices{
		ID:      uuid.New(),
		OrderID: idTrx,
		Email:   request.Email,
		Amount:  totalTrx,
		Status:  "Pending",
	}).Error
	if err != nil {
		tx.Rollback()
		return errors.New("failed to create invoices")
	}

	if request.PaymentMethod == "" || request.PaymentMethod == "cash" {
		fmt.Println("Payment method is cash, do not send email")
	}

	if err := r.db.Model(&entity.Transaction{}).Where("id = ? AND company_id = ?", idTrx, userInfo.CompanyID).Update("total", totalTrx).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) InputTransactionEach(request transactionDto.RequestTransactionEach, userInfo *middlewareDto.UserInfo) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	idTrx := uuid.New()
	orderDate := time.Now().Format("2006-01-02 15:04:05")
	totalTrx := 0

	// Create Transaction
	err := r.db.Create(&entity.Transaction{
		ID:        idTrx,
		UserID:    uuid.MustParse(userInfo.ID),
		OrderDate: orderDate,
		Customer:  request.Customer,
		Status:    false,
		Total:     totalTrx,
		CompanyID: uuid.MustParse(userInfo.CompanyID),
	}).Error
	if err != nil {
		tx.Rollback()
	}

	// Get Product Data, Create Detail Transaction and Invoices
	productRepo := product_repository.NewProductRepository(r.db)
	for _, v := range request.RequestProduct {
		product, err := productRepo.RetrieveByID(v.ProductID.String(), userInfo)
		if err != nil {
			tx.Rollback()
		}

		if !product.Status {
			tx.Rollback()
		}

		total := product.Price * v.Quantity

		// Create Detail Transaction
		err = r.db.Create(&entity.DetailTransaction{
			ID:        uuid.New(),
			OrderID:   idTrx,
			ProductID: product.ID,
			Quantity:  v.Quantity,
			Total:     total,
			CompanyID: product.CompanyID,
		}).Error
		if err != nil {
			tx.Rollback()
		}

		// Create Invoices
		err = r.db.Create(&entity.Invoices{
			ID:      uuid.New(),
			OrderID: idTrx,
			Email:   v.Email,
			Amount:  total,
			Status:  "Pending",
		}).Error
		if err != nil {
			tx.Rollback()
		}

		totalTrx += total
	}

	if err := r.db.Model(&entity.Transaction{}).Where("id = ? AND company_id = ?", idTrx, userInfo.CompanyID).Update("total", totalTrx).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) InputTransactionEqual(request transactionDto.RequestTransactionEqual, emailList []string, emailCount int, userInfo *middlewareDto.UserInfo) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	idTrx := uuid.New()
	orderDate := time.Now().Format("2006-01-02 15:04:05")
	totalTrx := 0

	// Create Transaction
	err := r.db.Create(&entity.Transaction{
		ID:        idTrx,
		UserID:    uuid.MustParse(userInfo.ID),
		OrderDate: orderDate,
		Customer:  request.Customer,
		Status:    false,
		Total:     totalTrx,
		CompanyID: uuid.MustParse(userInfo.CompanyID),
	}).Error
	if err != nil {
		tx.Rollback()
	}

	// Get Product Data, Create Detail Transaction and Invoices
	productRepo := product_repository.NewProductRepository(r.db)
	for _, v := range request.RequestProduct {
		product, err := productRepo.RetrieveByID(v.ProductID.String(), userInfo)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return errors.New("product with ID " + v.ProductID.String() + " not found")
			}
		}

		if !product.Status {
			tx.Rollback()
			return errors.New("product with ID " + v.ProductID.String() + " is not available")
		}

		total := product.Price * v.Quantity
		totalTrx += total

		// Create Detail Transaction
		err = r.db.Create(&entity.DetailTransaction{
			ID:        uuid.New(),
			OrderID:   idTrx,
			ProductID: product.ID,
			Quantity:  v.Quantity,
			Total:     total,
			CompanyID: product.CompanyID,
		}).Error
		if err != nil {
			tx.Rollback()
		}
	}

	amountPerEmail := int(math.Ceil(float64(totalTrx) / float64(emailCount)))

	// Create Invoices
	for _, email := range emailList {
		err = r.db.Create(&entity.Invoices{
			ID:      uuid.New(),
			OrderID: idTrx,
			Email:   email,
			Amount:  amountPerEmail,
			Status:  "Pending",
		}).Error
		if err != nil {
			tx.Rollback()
		}
	}

	if err := r.db.Model(&entity.Transaction{}).Where("id = ? AND company_id = ?", idTrx, userInfo.CompanyID).Update("total", totalTrx).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) RetrieveAllTransaction(page, size int, orderDate, status string, totalData int, userInfo *middlewareDto.UserInfo) (transactionDto.GetResponseTransaction, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size
	var transactions []entity.Transaction

	query := r.db.Where("company_id = ?", uuid.MustParse(userInfo.CompanyID))
	if status != "" {
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			return transactionDto.GetResponseTransaction{}, errors.New("invalid status value")
		}
		query = query.Where("status = ?", statusBool)
	}
	if orderDate != "" {
		query = query.Where("order_date = ?", orderDate)
	}

	if err := query.Limit(size).Offset(offset).Find(&transactions).Error; err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}
	fmt.Println(transactions)
	var responseTransactions []transactionDto.ResponseTransaction
	for _, trx := range transactions {
		var details []entity.DetailTransaction
		if err := r.db.Where("order_id = ?", trx.ID).Find(&details).Error; err != nil {
			return transactionDto.GetResponseTransaction{}, err
		}

		var products []transactionDto.GetProduct
		for _, detail := range details {
			var product productDto.Product
			if err := r.db.Where("id = ?", detail.ProductID).First(&product).Error; err != nil {
				return transactionDto.GetResponseTransaction{}, err
			}

			products = append(products, transactionDto.GetProduct{
				Name:     product.Name,
				Price:    product.Price,
				Quantity: detail.Quantity,
				Total:    detail.Total,
			})
		}

		var invoices []entity.Invoices
		if err := r.db.Where("order_id = ?", trx.ID).Find(&invoices).Error; err != nil {
			return transactionDto.GetResponseTransaction{}, err
		}

		var responseInvoices []transactionDto.GetInvoice
		for _, invoice := range invoices {
			responseInvoices = append(responseInvoices, transactionDto.GetInvoice{
				Email:  invoice.Email,
				Amount: invoice.Amount,
				Status: invoice.Status,
			})
		}

		var user userDto.User
		if err := r.db.Where("id = ?", trx.UserID).First(&user).Error; err != nil {
			return transactionDto.GetResponseTransaction{}, err
		}

		responseTransactions = append(responseTransactions, transactionDto.ResponseTransaction{
			ID:        trx.ID,
			OrderDate: trx.OrderDate,
			Customer:  trx.Customer,
			Status:    trx.Status,
			Total:     trx.Total,
			Employee: transactionDto.GetEmployee{
				ID:   user.ID,
				Name: user.Name,
			},
			Products: products,
			Invoices: responseInvoices,
		})
	}

	if len(responseTransactions) == 0 {
		return transactionDto.GetResponseTransaction{}, errors.New("data not found")
	}

	getResponseTransaction := transactionDto.GetResponseTransaction{
		Transactions: responseTransactions,
		Paging: transactionDto.Paging{
			Page:  page,
			Limit: size,
		},
		TotalData: totalData,
	}

	return getResponseTransaction, nil
}

func (r *transactionRepository) CountAllTransaction(page, size int, orderDate, status string, userInfo *middlewareDto.UserInfo) (int, error) {
	var count int64
	query := r.db.Model(&entity.Transaction{}).Where("company_id = ?", userInfo.CompanyID)
	if status != "" {
		statusBool, err := strconv.ParseBool(status)
		if err != nil {
			return 0, errors.New("invalid status value")
		}
		query = query.Where("status = ?", statusBool)
	}
	if orderDate != "" {
		query = query.Where("order_date = ?", orderDate)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *transactionRepository) RetrieveTransactionByID(id string, userInfo *middlewareDto.UserInfo) (transactionDto.GetResponseTransaction, error) {
	var transaction entity.Transaction
	if err := r.db.Where("id = ? AND company_id = ?", id, userInfo.CompanyID).First(&transaction).Error; err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	var details []entity.DetailTransaction
	if err := r.db.Where("order_id = ?", transaction.ID).Find(&details).Error; err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	var products []transactionDto.GetProduct
	for _, detail := range details {
		var product productDto.Product
		if err := r.db.Where("id = ?", detail.ProductID).First(&product).Error; err != nil {
			return transactionDto.GetResponseTransaction{}, err
		}

		products = append(products, transactionDto.GetProduct{
			Name:     product.Name,
			Price:    product.Price,
			Quantity: detail.Quantity,
			Total:    detail.Total,
		})
	}

	var invoices []entity.Invoices
	if err := r.db.Where("order_id = ?", transaction.ID).Find(&invoices).Error; err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	var responseInvoices []transactionDto.GetInvoice
	for _, invoice := range invoices {
		responseInvoices = append(responseInvoices, transactionDto.GetInvoice{
			Email:  invoice.Email,
			Amount: invoice.Amount,
			Status: invoice.Status,
		})
	}

	var user userDto.User
	if err := r.db.Where("id = ?", transaction.UserID).First(&user).Error; err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	responseTransaction := transactionDto.ResponseTransaction{
		ID:        transaction.ID,
		OrderDate: transaction.OrderDate,
		Customer:  transaction.Customer,
		Status:    transaction.Status,
		Total:     transaction.Total,
		Employee: transactionDto.GetEmployee{
			ID:   user.ID,
			Name: user.Name,
		},
		Products: products,
		Invoices: responseInvoices,
	}

	getResponseTransaction := transactionDto.GetResponseTransaction{
		Transactions: []transactionDto.ResponseTransaction{responseTransaction},
		TotalData:    1,
		Paging: transactionDto.Paging{
			Page:  1,
			Limit: 1,
		},
	}

	return getResponseTransaction, nil
}
