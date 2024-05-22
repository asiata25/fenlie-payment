package productDelivery

import (
	"finpro-fenlie/model/dto/json"
	"finpro-fenlie/model/entity"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productDelivery struct {
	productUC product.ProductUsecase
}

func NewProductDelivery(v1Group *gin.RouterGroup, productUC product.ProductUsecase) {
	handler := productDelivery{
		productUC: productUC,
	}
	productGroup := v1Group.Group("products")
	{
		productGroup.GET("", handler.GetAllProducts)
		productGroup.POST("", handler.CreateProduct)
		productGroup.GET("/:id", handler.GetProduct)
		productGroup.PUT("/:id", handler.UpdateProduct)
		productGroup.DELETE("/:id", handler.DeleteProduct)
	}
}

func (c *productDelivery) GetAllProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	products, _, _, totalItems, err := c.productUC.GetAllProducts(page, pageSize)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseWithPaging(ctx, products, page, int(totalItems))
}

func (c *productDelivery) CreateProduct(ctx *gin.Context) {
	var product entity.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequest(ctx, validationError, "bad request")
			return
		}
	}

	createdProduct, err := c.productUC.CreateProduct(product)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, createdProduct, "success sreate product")
}

func (c *productDelivery) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product entity.Product

	getProduct, err := c.productUC.GetProduct(id, product)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, getProduct, "uccess create product")
}

func (c *productDelivery) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product entity.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequest(ctx, validationError, "bad request")
			return
		}
	}

	_, err := c.productUC.UpdateProduct(id, product)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	result, err := c.productUC.GetProduct(id, product)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, result, "Success Update Product")
}

func (c *productDelivery) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.productUC.DeleteProduct(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "SuccessFuly Delete Product")
}
