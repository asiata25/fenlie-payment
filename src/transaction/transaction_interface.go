package transaction

import (
	"finpro-fenlie/model/dto/transaction"
	"finpro-fenlie/model/entity"
)

type TransactionRepository interface {
	InputTransaction(payload entity.Transaction) error
	// InputTransactionEach(transaction.RequestTransaction, *middlewareDto.UserInfo) error
	// InputTransactionEqual(transaction.RequestTransaction, []string, int, *middlewareDto.UserInfo) error
	RetrieveAllTransaction(page, size int, orderDate, status, companyId string) ([]entity.Transaction, int, error)
	RetrieveTransactionByID(id, companyId string) (entity.Transaction, error)
	EditTransaction(id, companyId string, payload map[string]interface{}) error
}

type TransactionUsecase interface {
	CreateTransaction(request transaction.RequestTransaction) error
	GetAllTransaction(page, size int, orderDate, status, companyId string) ([]transaction.ResponseTransaction, int, error)
	GetTransactionByID(id, companyId string) (transaction.ResponseTransaction, error)
	UpdateTransaction(id, companyId string, payload map[string]interface{}) error
	// CreateTransactionEach(*gin.Context, transactionDto.RequestTransactionEach) error
	// CreateTransactionEqual(*gin.Context, transactionDto.RequestTransactionEqual) error

	// ReadAllTransaction(*gin.Context, int, int, string, string) (transactionDto.GetResponseTransaction, error)

	// ReadTransactionByID(*gin.Context, string) (transactionDto.GetResponseTransaction, error)
}
