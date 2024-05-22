package categoryDelivery

import (
	categoryDto "finpro-fenlie/model/dto/category"
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
	var request categoryDto.CategoryRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseBadRequest(ctx, validationErr, "bad request")
			return
		}
		jsonDTO.NewResponseError(ctx, "no request body found")
		return
	}

	err := cd.useCase.CreateLoan(&request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, request, "created successfully")
}

func (cd *categoryDelivery) GetAll(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")

	loans, total, err := cd.useCase.GetAllLoans(page, size)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseWithPaging(ctx, loans, pageInt, total)
}

func (cd *categoryDelivery) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	loans, err := cd.useCase.GetLoanById(ID)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, loans, "ok")
}

func (cd *categoryDelivery) Update(ctx *gin.Context) {
	ID := ctx.Param("id")
	var request categoryDto.CategoryRequest

	request.ID = ID
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseBadRequest(ctx, validationErr, "bad request")
			return
		}
		jsonDTO.NewResponseError(ctx, "no request body found")
		return
	}

	err := cd.useCase.UpdateLoan(&request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "updated successfully")
}

func (cd *categoryDelivery) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	err := cd.useCase.DeleteLoan(ID)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "delete success")
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
