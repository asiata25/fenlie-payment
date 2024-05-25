package transactionUsecase

import (
	"errors"
	"finpro-fenlie/model/dto/transactionDto"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/src/transaction"
	"strings"

	"github.com/gin-gonic/gin"
)

type transactionUC struct {
	transactionRepo transaction.TransactionRepository
}

func NewTransactionUsecase(transactionRepo transaction.TransactionRepository) transaction.TransactionUsecase {
	return &transactionUC{transactionRepo}
}

// Implement CreateTransaction
func (usecase *transactionUC) CreateTransaction(c *gin.Context, request transactionDto.RequestTransaction) error {
	userInfo, err := middleware.GetUserInfo(c)
	if err != nil {
		return err
	}

	err = usecase.transactionRepo.InputTransaction(request, userInfo)
	if err != nil {
		return err
	}

	return nil

}

// Implement CreateTransactionEach
func (usecase *transactionUC) CreateTransactionEach(c *gin.Context, request transactionDto.RequestTransactionEach) error {
	userInfo, err := middleware.GetUserInfo(c)
	if err != nil {
		return err
	}

	err = usecase.transactionRepo.InputTransactionEach(request, userInfo)
	if err != nil {
		return err
	}

	return nil
}

// Implement CreateTransactionEqual
func (usecase *transactionUC) CreateTransactionEqual(c *gin.Context, req transactionDto.RequestTransactionEqual) error {
	userInfo, err := middleware.GetUserInfo(c)
	if err != nil {
		return err
	}

	// Split and Count email
	emailList := strings.Split(req.Email, ", ")
	emailCount := len(emailList)
	if emailCount == 0 {
		return errors.New("email cannot be empty")
	}
	for i := range emailList {
		emailList[i] = strings.TrimSpace(emailList[i])
	}

	err = usecase.transactionRepo.InputTransactionEqual(req, emailList, emailCount, userInfo)
	if err != nil {
		return err
	}

	return nil
}

// Implement ReadAllTransaction
func (usecase *transactionUC) ReadAllTransaction(c *gin.Context, page, size int, orderDate, status string) (transactionDto.GetResponseTransaction, error) {
	userInfo, err := middleware.GetUserInfo(c)
	if err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	totalData, err := usecase.transactionRepo.CountAllTransaction(page, size, orderDate, status, userInfo)
	if err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	if totalData == 0 {
		return transactionDto.GetResponseTransaction{}, errors.New("data not found")
	}

	transactions, err := usecase.transactionRepo.RetrieveAllTransaction(page, size, orderDate, status, totalData, userInfo)
	if err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	return transactions, nil
}

// Implement ReadTransactionByID
func (usecase *transactionUC) ReadTransactionByID(c *gin.Context, id string) (transactionDto.GetResponseTransaction, error) {
	userInfo, err := middleware.GetUserInfo(c)
	if err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	transactions, err := usecase.transactionRepo.RetrieveTransactionByID(id, userInfo)
	if err != nil {
		return transactionDto.GetResponseTransaction{}, err
	}

	return transactions, nil
}
