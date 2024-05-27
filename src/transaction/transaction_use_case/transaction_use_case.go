package transactionUsecase

import (
	transactionDTO "finpro-fenlie/model/dto/transaction"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/pkg/email"
	"finpro-fenlie/src/transaction"
)

type transactionUC struct {
	transactionRepo transaction.TransactionRepository
}

func NewTransactionUsecase(transactionRepo transaction.TransactionRepository) transaction.TransactionUsecase {
	return &transactionUC{transactionRepo}
}

// Implement CreateTransaction
func (usecase *transactionUC) CreateTransaction(request transactionDTO.RequestTransaction) error {
	var detailTransactions []entity.DetailTransaction
	for _, detail := range request.DetailTransactions {
		detailTransactions = append(detailTransactions, entity.DetailTransaction{
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Amount:    detail.Amount,
		})
	}

	var invoices []entity.Invoice
	for _, invoice := range request.Invoices {
		invoices = append(invoices, entity.Invoice{
			EmailCustomer: invoice.EmailCustomer,
			Amount:        invoice.Amount,
			Status:        "unpaid",
		})
	}

	transaction := entity.Transaction{
		Status:             "unpaid",
		CompanyID:          request.CompanyID,
		UserId:             request.UserID,
		DetailTransactions: detailTransactions,
		Invoices:           invoices,
		Total:              request.Total,
	}
	err := usecase.transactionRepo.InputTransaction(transaction)
	if err != nil {
		return err
	}

	for _, invoice := range invoices {
		err = email.Send(invoice.EmailCustomer, "Invoice Created", "body")
		if err != nil {
			return err
		}
	}

	return nil

}

// // Implement CreateTransactionEach
// func (usecase *transactionUC) CreateTransactionEach(c *gin.Context, request transactionDto.RequestTransactionEach) error {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return err
// 	}

// 	err = usecase.transactionRepo.InputTransactionEach(request, userInfo)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Implement CreateTransactionEqual
// func (usecase *transactionUC) CreateTransactionEqual(c *gin.Context, req transactionDto.RequestTransactionEqual) error {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return err
// 	}

// 	// Split and Count email
// 	emailList := strings.Split(req.Email, ", ")
// 	emailCount := len(emailList)
// 	if emailCount == 0 {
// 		return errors.New("email cannot be empty")
// 	}
// 	for i := range emailList {
// 		emailList[i] = strings.TrimSpace(emailList[i])
// 	}

// 	err = usecase.transactionRepo.InputTransactionEqual(req, emailList, emailCount, userInfo)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Implement ReadAllTransaction
// func (usecase *transactionUC) ReadAllTransaction(c *gin.Context, page, size int, orderDate, status string) (transactionDto.GetResponseTransaction, error) {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return transactionDto.GetResponseTransaction{}, err
// 	}

// 	totalData, err := usecase.transactionRepo.CountAllTransaction(page, size, orderDate, status, userInfo)
// 	if err != nil {
// 		return transactionDto.GetResponseTransaction{}, err
// 	}

// 	if totalData == 0 {
// 		return transactionDto.GetResponseTransaction{}, errors.New("data not found")
// 	}

// 	transactions, err := usecase.transactionRepo.RetrieveAllTransaction(page, size, orderDate, status, totalData, userInfo)
// 	if err != nil {
// 		return transactionDto.GetResponseTransaction{}, err
// 	}

// 	return transactions, nil
// }

// // Implement ReadTransactionByID
// func (usecase *transactionUC) ReadTransactionByID(c *gin.Context, id string) (transactionDto.GetResponseTransaction, error) {
// 	userInfo, err := middleware.GetUserInfo(c)
// 	if err != nil {
// 		return transactionDto.GetResponseTransaction{}, err
// 	}

// 	transactions, err := usecase.transactionRepo.RetrieveTransactionByID(id, userInfo)
// 	if err != nil {
// 		return transactionDto.GetResponseTransaction{}, err
// 	}

// 	return transactions, nil
// }
