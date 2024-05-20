package categoryDelivery

import (
	"finpro-fenlie/model/dto/categoryDto"
	jsonDTO "finpro-fenlie/model/dto/json"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/category"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryDelivery struct {
	useCase category.CategoryUseCase
}

func (cd *categoryDelivery) Create(ctx *gin.Context) {
	var request categoryDto.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseBadRequest(ctx, validationErr, "bad request", "01", "01")
			return
		}
		jsonDTO.NewResponseError(ctx, "no request body found", "01", "02")
		return
	}

	err := cd.useCase.CreateLoan(&request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	jsonDTO.NewResponseSuccess(ctx, request, "created successfully", "01", "01")
}

func (cd *categoryDelivery) GetAll(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")

	loans, total, err := cd.useCase.GetAllLoans(page, size)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	jsonDTO.NewResponseUserPaging(ctx, loans, pageInt, total, "01", "01")
}

func (cd *categoryDelivery) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	loans, err := cd.useCase.GetLoanById(ID)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	jsonDTO.NewResponseSuccess(ctx, loans, "ok", "01", "01")
}

func (cd *categoryDelivery) Update(ctx *gin.Context) {
	ID := ctx.Param("id")
	var request categoryDto.UpdateCategoryRequest

	request.ID = ID
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseBadRequest(ctx, validationErr, "bad request", "01", "01")
			return
		}
		jsonDTO.NewResponseError(ctx, "no request body found", "01", "02")
		return
	}

	err := cd.useCase.UpdateLoan(&request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "updated successfully", "01", "01")
}

func (cd *categoryDelivery) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := cd.useCase.DeleteLoan(ID)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "delete success", "01", "01")
}

func NewCategoryDelivery(v1Group *gin.RouterGroup, useCase category.CategoryUseCase) {
	handler := categoryDelivery{useCase}

	loans := v1Group.Group("/category")
	{
		loans.POST("", handler.Create)
		loans.GET("", handler.GetAll)
		loans.GET("/:id", handler.GetByID)
		loans.PUT("/:id", handler.Update)
		loans.DELETE("/:id", handler.Delete)
	}
}
