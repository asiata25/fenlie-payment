package userDelivery

import (
	"finpro-fenlie/model/dto/json"
	userDTO "finpro-fenlie/model/dto/user"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/pkg/validation"
	"finpro-fenlie/src/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUC user.UserUseCase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC user.UserUseCase) {
	handler := UserDelivery{userUC}
	users := v1Group.Group("/users")
	{
		users.POST("/login", handler.login)
		users.POST("", middleware.JWTAuth("ADMIN"), handler.createUser)
		users.GET("", middleware.JWTAuth("ADMIN"), handler.getUser)
		users.GET("/:id", middleware.JWTAuth("ADMIN"), handler.getUserById)
		users.PUT("/:id", middleware.JWTAuth("ADMIN"), handler.updateUser)
		users.DELETE("/:id", middleware.JWTAuth("ADMIN"), handler.deleteUser)

	}
}

func (u *UserDelivery) login(ctx *gin.Context) {
	var req userDTO.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	companyId := ctx.GetHeader("companyId")
	req.CompanyID = companyId

	token, err := u.userUC.Login(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, gin.H{"token": token}, "success")
}

func (u *UserDelivery) createUser(ctx *gin.Context) {
	var user userDTO.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			json.NewResponseValidationError(ctx, validationErr, "bad request")
			return
		}

		json.NewResponseError(ctx, "no request body found")
		return
	}

	companyId := ctx.GetHeader("companyId")
	user.CompanyID = companyId

	err := u.userUC.CreateUser(user)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "success")
}

func (u *UserDelivery) getUser(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 0 {
		json.NewResponseBadRequest(ctx, "fail to convert data", "page is not a positive number")
		return
	}

	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil || size < 0 {
		json.NewResponseBadRequest(ctx, "fail to convert data", "size is not a positive number")
		return
	}

	email := ctx.Query("email")
	name := ctx.Query("name")

	companyId := ctx.GetHeader("companyId")
	users, total, err := u.userUC.GetAllUser(page, size, email, name, companyId)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseWithPaging(ctx, users, page, total)
}

func (u *UserDelivery) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")

	users, err := u.userUC.GetUserByID(id, companyId)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, users, "success")
}

func (u *UserDelivery) updateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")

	var request userDTO.UpdateUserRequest
	request.ID = id
	request.CompanyID = companyId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationErr := validation.GetValidationError(err)
		if len(validationErr) > 0 {
			json.NewResponseValidationError(ctx, validationErr, "bad request")
			return
		}

		json.NewResponseError(ctx, "no request body found")
		return
	}

	err := u.userUC.UpdateUser(request)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "success")
}

func (u *UserDelivery) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	companyId := ctx.GetHeader("companyId")
	currentUserId := ctx.GetHeader("userId")

	if currentUserId == id {
		json.NewResponseError(ctx, "can't delete your own account")
		return
	}

	err := u.userUC.DeleteUser(id, companyId)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, nil, "success")
}
