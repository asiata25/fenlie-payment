package transactionDelivery

import (
	jsonDto "finpro-fenlie/model/dto/json"
	"finpro-fenlie/model/dto/transactionDto"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/src/transaction"
	"strconv"

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
	transactions.Use(middleware.JWTAuth(), middleware.UserOnly())
	{
		transactions.POST("", handler.createTransaction)
		transactions.POST("/split-each", handler.createTransactionEach)
		transactions.POST("/split-equal", handler.createTransactionEqual)
		transactions.GET("", handler.getAllTransaction)
		transactions.GET("/:id", handler.getTransactionByID)
	}
}

func (t *TransactionDelivery) createTransaction(ctx *gin.Context) {
	// Bind request
	var req transactionDto.RequestTransaction
	if err := ctx.ShouldBindJSON(&req); err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	// Create transaction
	err := t.transactionUC.CreateTransaction(ctx, req)
	if err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	jsonDto.NewResponseSuccess(ctx, nil, "success")
}

func (t *TransactionDelivery) createTransactionEach(ctx *gin.Context) {
	// Bind request
	var req transactionDto.RequestTransactionEach
	if err := ctx.ShouldBindJSON(&req); err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	// Create transaction
	err := t.transactionUC.CreateTransactionEach(ctx, req)
	if err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	jsonDto.NewResponseSuccess(ctx, nil, "success")
}

func (t *TransactionDelivery) createTransactionEqual(ctx *gin.Context) {
	// Bind request
	var req transactionDto.RequestTransactionEqual
	if err := ctx.ShouldBindJSON(&req); err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	// Create transaction
	err := t.transactionUC.CreateTransactionEqual(ctx, req)
	if err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	jsonDto.NewResponseSuccess(ctx, nil, "success")
}

func (t *TransactionDelivery) getAllTransaction(ctx *gin.Context) {
	// Get query
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	orderDate := ctx.Query("orderDate")
	status := ctx.Query("status")

	// Get all transaction
	res, err := t.transactionUC.ReadAllTransaction(ctx, page, size, orderDate, status)
	if err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	jsonDto.NewResponseSuccess(ctx, res, "success")
}

func (t *TransactionDelivery) getTransactionByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Get transaction by ID
	res, err := t.transactionUC.ReadTransactionByID(ctx, id)
	if err != nil {
		jsonDto.NewResponseError(ctx, err.Error())
		return
	}

	jsonDto.NewResponseSuccess(ctx, res, "success")
}
