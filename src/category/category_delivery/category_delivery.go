package categoryDelivery

import (
	categoryDto "finpro-fenlie/model/dto/category"
	jsonDTO "finpro-fenlie/model/dto/json"
	"finpro-fenlie/pkg/middleware"
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
			jsonDTO.NewResponseValidationError(ctx, validationErr, "bad request")
			return
		}
		jsonDTO.NewResponseError(ctx, "no request body found")
		return
	}

	companyId := ctx.GetHeader("companyId")
	request.CompanyID = companyId

	err := cd.useCase.Create(&request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "created successfully")
}

func (cd *categoryDelivery) GetAll(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}
	name := ctx.Query("name")

	companyId := ctx.GetHeader("companyId")
	categories, total, err := cd.useCase.GetAll(page, size, name, companyId)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseWithPaging(ctx, categories, page, total)
}

func (cd *categoryDelivery) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")
	loans, err := cd.useCase.GetById(ID, companyId)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, loans, "ok")
}

func (cd *categoryDelivery) Update(ctx *gin.Context) {

	var request categoryDto.CategoryRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseValidationError(ctx, validationErr, "bad request")
			return
		}
		jsonDTO.NewResponseError(ctx, "no request body found")
		return
	}

	ID := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")

	request.ID = ID
	request.CompanyID = companyId

	err := cd.useCase.Update(&request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "updated successfully")
}

func (cd *categoryDelivery) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")

	err := cd.useCase.Delete(ID, companyId)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "delete success")
}

func NewCategoryDelivery(v1Group *gin.RouterGroup, useCase category.CategoryUseCase) {
	handler := categoryDelivery{useCase}

	category := v1Group.Group("/categories")
	category.Use(middleware.JWTAuth("ADMIN", "EMPLOYEE"))
	{
		category.POST("", handler.Create)
		category.GET("", handler.GetAll)
		category.GET("/:id", handler.GetByID)
		category.PUT("/:id", handler.Update)
		category.DELETE("/:id", handler.Delete)
	}
}
