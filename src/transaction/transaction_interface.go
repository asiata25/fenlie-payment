package transaction

import (
	"finpro-fenlie/model/dto/transaction"
	"finpro-fenlie/model/entity"
)

type TransactionRepository interface {
	InputTransaction(payload entity.Transaction) error
	// InputTransactionEach(transactionDto.RequestTransactionEach, *middlewareDto.UserInfo) error
	// InputTransactionEqual(transactionDto.RequestTransactionEqual, []string, int, *middlewareDto.UserInfo) error

	// RetrieveAllTransaction(int, int, string, string, int, *middlewareDto.UserInfo) (transactionDto.GetResponseTransaction, error)
	// CountAllTransaction(int, int, string, string, *middlewareDto.UserInfo) (int, error)

	// RetrieveTransactionByID(string, *middlewareDto.UserInfo) (transactionDto.GetResponseTransaction, error)
}

type TransactionUsecase interface {
	CreateTransaction(request transaction.RequestTransaction) error
	// CreateTransactionEach(*gin.Context, transactionDto.RequestTransactionEach) error
	// CreateTransactionEqual(*gin.Context, transactionDto.RequestTransactionEqual) error

	// ReadAllTransaction(*gin.Context, int, int, string, string) (transactionDto.GetResponseTransaction, error)

	// ReadTransactionByID(*gin.Context, string) (transactionDto.GetResponseTransaction, error)
}
