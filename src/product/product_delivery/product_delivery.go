package productDelivery

import (
	"finpro-fenlie/model/dto/json"
	productDTO "finpro-fenlie/model/dto/product"
	cloudx "finpro-fenlie/pkg/cloudinary"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/product"
	"fmt"
	"log"
	"net/http"
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
	productGroup.Use(middleware.JWTAuth("ADMIN", "EMPLOYEE"))
	{
		productGroup.GET("", handler.GetAllProducts)
		productGroup.POST("", handler.CreateProduct)
		productGroup.GET("/:id", handler.GetProduct)
		productGroup.PUT("/:id", handler.UpdateProduct)
		productGroup.DELETE("/:id", handler.DeleteProduct)
	}
}

func (c *productDelivery) GetAllProducts(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	name := ctx.Query("name")
	companyId := ctx.GetHeader("companyId")
	products, totalItems, err := c.productUC.GetAllProducts(page, pageSize, name, companyId)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseWithPaging(ctx, products, page, int(totalItems))
}

func (c *productDelivery) CreateProduct(ctx *gin.Context) {

	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Failed to parse form data: %v", err)
		json.NewResponseBadRequest(ctx, nil, "failed to parse form data")
		return
	}

	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		log.Printf("Failed to get image: %v", err)
		json.NewResponseBadRequest(ctx, nil, "failed to get image")
		return
	}
	defer file.Close()

	name := ctx.Request.FormValue("name")
	priceStr := ctx.Request.FormValue("price")
	statusStr := ctx.Request.FormValue("status")

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		log.Printf("Invalid price value: %v", err)
		http.Error(ctx.Writer, "Invalid price value", http.StatusBadRequest)
		return
	}

	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		log.Printf("Invalid status value: %v", err)
		http.Error(ctx.Writer, "Invalid status value", http.StatusBadRequest)
		return
	}

	companyId := ctx.GetHeader("companyId")

	product := productDTO.ProductCreateRequest{
		Name:      name,
		Price:     price,
		Status:    status,
		CompanyID: companyId,
	}

	imageURL, err := cloudx.UploadImage(file)
	if err != nil {
		log.Printf("Failed to upload image: %v", err)
		http.Error(ctx.Writer, "Failed to upload image", http.StatusInternalServerError)
		return
	}
	fmt.Println(imageURL)
	err = c.productUC.CreateProduct(product, imageURL)
	if err != nil {
		log.Printf("Failed to create product: %v", err)
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "success create product")
}

func (c *productDelivery) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")
	getProduct, err := c.productUC.GetProduct(id, companyId)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, getProduct, "OK")
}

func (c *productDelivery) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")
	var product productDTO.ProductUpdateRequest
	product.ID = id
	product.CompanyID = companyId

	if err := ctx.ShouldBindJSON(&product); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequest(ctx, validationError, "bad request")
			return
		}
	}

	err := c.productUC.UpdateProduct(product)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "Success Update Product")
}

func (c *productDelivery) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")

	err := c.productUC.DeleteProduct(id, companyId)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "SuccessFuly Delete Product")
}
