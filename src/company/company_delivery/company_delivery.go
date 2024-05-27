package companyDelivery

import (
	companyDTO "finpro-fenlie/model/dto/company"
	jsonDTO "finpro-fenlie/model/dto/json"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/company"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type companyDelivery struct {
	useCase company.CompanyUseCase
}

func (c *companyDelivery) getById(ctx *gin.Context) {
	companyId := ctx.Param("companyId")

	company, err := c.useCase.GetById(companyId)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, company, "success")
}

func (c *companyDelivery) create(ctx *gin.Context) {
	var request companyDTO.CompanyCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseValidationError(ctx, validationErr, "bad request")
			return
		}

		jsonDTO.NewResponseError(ctx, "no request body found")
		return
	}

	err := c.useCase.Create(request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "success")
}

func (c *companyDelivery) delete(ctx *gin.Context) {
	companyId := ctx.Param("companyId")

	err := c.useCase.Delete(companyId)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "success delete company")
}

func (c *companyDelivery) getAll(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 0 {
		jsonDTO.NewResponseBadRequest(ctx, "fail to convert data", "page is not a positive number")
		return
	}

	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil || size < 0 {
		jsonDTO.NewResponseBadRequest(ctx, "fail to convert data", "size is not a positive number")
		return
	}
	name := ctx.Query("name")

	companies, total, err := c.useCase.GetAll(page, size, name)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseWithPaging(ctx, companies, page, int(total))
}

func (c *companyDelivery) update(ctx *gin.Context) {
	var request companyDTO.CompanyUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			jsonDTO.NewResponseValidationError(ctx, validationErr, "bad request")
			return
		}

		jsonDTO.NewResponseError(ctx, "no request body found")
		return
	}

	companyId := ctx.Param("companyId")
	request.ID = companyId

	err := c.useCase.Update(request)
	if err != nil {
		jsonDTO.NewResponseError(ctx, err.Error())
		return
	}

	jsonDTO.NewResponseSuccess(ctx, nil, "success update company")
}

func NewCompanyDelivery(v1Group *gin.RouterGroup, useCase company.CompanyUseCase) {
	handler := companyDelivery{useCase}

	company := v1Group.Group("/companies")
	company.Use(gin.BasicAuth(gin.Accounts{
		os.Getenv("CLIENT_ID"): os.Getenv("CLIENT_SECRET"),
	}))
	{
		company.GET("/:companyId", handler.getById)
		company.POST("", handler.create)
		company.DELETE("/:companyId", handler.delete)
		company.GET("", handler.getAll)
		company.PUT("/:companyId", handler.update)
	}
}
