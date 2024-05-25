package transactionDelivery

import (
	jsonDTO "finpro-fenlie/model/dto/json"
	transactionDTO "finpro-fenlie/model/dto/transaction"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/transaction"

	"github.com/gin-gonic/gin"
)

type TransactionDelivery struct {
	transactionUC transaction.TransactionUsecase
}

func NewTransactionDelivery(v1Group *gin.RouterGroup, transactionUC transaction.TransactionUsecase) {
	handler := TransactionDelivery{
		transactionUC: transactionUC,
	}
	transactions := v1Group.Group("/transactions")
	transactions.Use(middleware.JWTAuth("ADMIN", "EMPLOYEE"))
	{
		transactions.POST("", handler.createTransaction)
		// transactions.POST("/split-each", handler.createTransactionEach)
		// transactions.POST("/split-equal", handler.createTransactionEqual)
		// transactions.GET("", handler.getAllTransaction)
		// transactions.GET("/:id", handler.getTransactionByID)
	}
}

func (t *TransactionDelivery) createTransaction(ctx *gin.Context) {
	// Bind request
	var req transactionDTO.RequestTransaction
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			jsonDTO.NewResponseBadRequest(ctx, validationError, "bad request")
			return
		}
	}
	companyId := ctx.GetHeader("companyId")
	currentUserId := ctx.GetHeader("userId")
	req.CompanyID = companyId
	req.UserID = currentUserId

	// Create transaction
	err := t.transactionUC.CreateTransaction(req)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "success")
}

// func (t *TransactionDelivery) createTransactionEach(ctx *gin.Context) {
// 	// Bind request
// 	var req transactionDto.RequestTransactionEach
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		jsonDto.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	// Create transaction
// 	err := t.transactionUC.CreateTransactionEach(ctx, req)
// 	if err != nil {
// 		jsonDto.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDto.NewResponseSuccess(ctx, nil, "success")
// }

// func (t *TransactionDelivery) createTransactionEqual(ctx *gin.Context) {
// 	// Bind request
// 	var req transactionDto.RequestTransactionEqual
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		jsonDto.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	// Create transaction
// 	err := t.transactionUC.CreateTransactionEqual(ctx, req)
// 	if err != nil {
// 		jsonDto.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDto.NewResponseSuccess(ctx, nil, "success")
// }

// func (t *TransactionDelivery) getAllTransaction(ctx *gin.Context) {
// 	// Get query
// 	page, _ := strconv.Atoi(ctx.Query("page"))
// 	size, _ := strconv.Atoi(ctx.Query("size"))
// 	orderDate := ctx.Query("orderDate")
// 	status := ctx.Query("status")

// 	// Get all transaction
// 	res, err := t.transactionUC.ReadAllTransaction(ctx, page, size, orderDate, status)
// 	if err != nil {
// 		jsonDto.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDto.NewResponseSuccess(ctx, res, "success")
// }

// func (t *TransactionDelivery) getTransactionByID(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	// Get transaction by ID
// 	res, err := t.transactionUC.ReadTransactionByID(ctx, id)
// 	if err != nil {
// 		jsonDto.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDto.NewResponseSuccess(ctx, res, "success")
// }
