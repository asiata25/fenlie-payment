package userDelivery

import (
	"finpro-fenlie/model/dto/json"
	"finpro-fenlie/model/dto/middlewareDto"
	"finpro-fenlie/model/dto/userDto"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/src/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUC user.UserUsecase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC user.UserUsecase) {
	handler := UserDelivery{
		userUC: userUC}
	users := v1Group.Group("/users")
	users.Use(middleware.JWTAuth(), middleware.AdminOnly())
	{
		users.POST("", handler.createUser)
		users.GET("", handler.getUser)
		users.GET("/:id", handler.getUserById)
		users.PUT("/:id", handler.updateUser)
		users.DELETE("/:id", handler.deleteUser)

	}
	login := v1Group.Group("/users/login")
	login.Use(middleware.BasicAuth)
	{
		login.POST("", handler.login)
	}
}

func (u *UserDelivery) login(ctx *gin.Context) {
	var req middlewareDto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	users, err := u.userUC.Login(ctx, req)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccessToken(ctx, users, "success", "01", "01")
}

func (u *UserDelivery) createUser(ctx *gin.Context) {
	var user userDto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	err := u.userUC.CreateUser(ctx, user)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "success", "01", "01")
}

func (u *UserDelivery) getUser(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	email := ctx.Query("email")
	name := ctx.Query("name")

	users, err := u.userUC.GetAllUser(ctx, page, size, email, name)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, users, "success", "01", "01")
}

func (u *UserDelivery) getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	users, err := u.userUC.GetUserByID(ctx, id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	} else {
		json.NewResponseSuccess(ctx, users, "success", "01", "01")
	}
}

func (u *UserDelivery) updateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var userUpdates map[string]interface{}
	if err := ctx.ShouldBindJSON(&userUpdates); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	err := u.userUC.UpdateUser(ctx, id, userUpdates)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "success", "01", "01")
}

func (u *UserDelivery) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := u.userUC.DeleteUser(ctx, id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "success", "01", "01")
}