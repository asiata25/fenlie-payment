package transaction

import (
	"finpro-fenlie/model/dto/middlewareDto"
	"finpro-fenlie/model/dto/transactionDto"

	"github.com/gin-gonic/gin"
)

type TransactionRepository interface {
	InputTransaction(transactionDto.RequestTransaction, *middlewareDto.UserInfo) error
	InputTransactionEach(transactionDto.RequestTransactionEach, *middlewareDto.UserInfo) error
	InputTransactionEqual(transactionDto.RequestTransactionEqual, []string, int, *middlewareDto.UserInfo) error

	RetrieveAllTransaction(int, int, string, string, int, *middlewareDto.UserInfo) (transactionDto.GetResponseTransaction, error)
	CountAllTransaction(int, int, string, string, *middlewareDto.UserInfo) (int, error)

	RetrieveTransactionByID(string, *middlewareDto.UserInfo) (transactionDto.GetResponseTransaction, error)
}

type TransactionUsecase interface {
	CreateTransaction(*gin.Context, transactionDto.RequestTransaction) error
	CreateTransactionEach(*gin.Context, transactionDto.RequestTransactionEach) error
	CreateTransactionEqual(*gin.Context, transactionDto.RequestTransactionEqual) error

	ReadAllTransaction(*gin.Context, int, int, string, string) (transactionDto.GetResponseTransaction, error)

	ReadTransactionByID(*gin.Context, string) (transactionDto.GetResponseTransaction, error)
}
