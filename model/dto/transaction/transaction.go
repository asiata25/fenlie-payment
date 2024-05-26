package transaction

import (
	"finpro-fenlie/model/entity"
	"time"
)

type (
	RequestTransaction struct {
		Total              int                 `json:"total" binding:"required,number"`
		CompanyID          string              `json:"company_id"`
		UserID             string              `json:"user_id"`
		DetailTransactions []DetailTransaction `json:"detail_transactions" binding:"required,dive"`
		Invoices           []Invoices          `json:"invoices" binding:"required,dive"`
	}

	DetailTransaction struct {
		ProductID string `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required,number"`
		Amount    int    `json:"amount" binding:"required,number"`
	}

	Invoices struct {
		EmailCustomer string `json:"email_customer" binding:"required,email"`
		Amount        int    `json:"amount" binding:"required,number"`
	}

	ResponseTransaction struct {
		ID                 string                     `json:"id"`
		OrderDate          time.Time                  `json:"order_date"`
		Status             string                     `json:"status"`
		CompanyID          string                     `json:"company_id"`
		UserID             string                     `json:"user_id"`
		DetailTransactions []entity.DetailTransaction `json:"detail_transactions"`
		Invoices           []entity.Invoice           `json:"invoices"`
		Total              int                        `json:"total"`
	}

	// RequestTransactionEach struct {
	// 	Customer       string               `json:"customer"`
	// 	RequestProduct []RequestProductEach `json:"request_product"`
	// }

	// RequestTransactionEqual struct {
	// 	Customer       string                `json:"customer"`
	// 	RequestProduct []RequestProductEqual `json:"request_product"`
	// 	Email          string                `json:"email"`
	// }

	// RequestProductEach struct {
	// 	ProductID string `json:"product_id"`
	// 	Quantity  int    `json:"quantity"`
	// 	Email     string `json:"email"`
	// }

	// RequestProductEqual struct {
	// 	ProductID string `json:"product_id"`
	// 	Quantity  int    `json:"quantity"`
	// }

	// GetEmployee struct {
	// 	ID   string `json:"id"`
	// 	Name string `json:"name"`
	// }

	// GetProduct struct {
	// 	Name     string `json:"name"`
	// 	Price    int    `json:"price"`
	// 	Quantity int    `json:"quantity"`
	// 	Total    int    `json:"total"`
	// }

	// GetInvoice struct {
	// 	Email  string `json:"email"`
	// 	Amount int    `json:"amount"`
	// 	Status string `json:"status"`
	// }

	// Paging struct {
	// 	Page  int `json:"page"`
	// 	Limit int `json:"limit"`
	// }

	// ResponseTransaction struct {
	// 	ID        string       `json:"id"`
	// 	OrderDate string       `json:"order_date"`
	// 	Customer  string       `json:"customer"`
	// 	Status    bool         `json:"status"`
	// 	Total     int          `json:"total"`
	// 	Employee  GetEmployee  `json:"employee"`
	// 	Products  []GetProduct `json:"products"`
	// 	Invoices  []GetInvoice `json:"invoices"`
	// }

	// GetResponseTransaction struct {
	// 	Transactions []ResponseTransaction `json:"transactions"`
	// 	TotalData    int                   `json:"total_data"`
	// 	Paging       Paging                `json:"paging"`
	// }
)
