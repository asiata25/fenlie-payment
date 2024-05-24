package userDelivery

import (
	"finpro-fenlie/model/dto/json"
	userDTO "finpro-fenlie/model/dto/user"
	"finpro-fenlie/pkg/middleware"
	"finpro-fenlie/src/user"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUC user.UserUseCase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC user.UserUseCase) {
	handler := UserDelivery{userUC}
	users := v1Group.Group("/users")
	users.Use(middleware.JWTAuth())
	{
		users.POST("/login", handler.login)
		// users.POST("", handler.createUser)
		// users.GET("", handler.getUser)
		// users.GET("/:id", handler.getUserById)
		// users.PUT("/:id", handler.updateUser)
		// users.DELETE("/:id", handler.deleteUser)

	}
}

func (u *UserDelivery) login(ctx *gin.Context) {
	var req userDTO.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	token, err := u.userUC.Login(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error())
		return
	}

	json.NewResponseSuccess(ctx, gin.H{"token": token}, "success")
}

// func (u *UserDelivery) createUser(ctx *gin.Context) {
// 	var user userDto.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	err := u.userUC.CreateUser(ctx, user)
// 	if err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDTO.NewResponseSuccess(ctx, nil, "success")
// }

// func (u *UserDelivery) getUser(ctx *gin.Context) {
// 	page, _ := strconv.Atoi(ctx.Query("page"))
// 	size, _ := strconv.Atoi(ctx.Query("size"))
// 	email := ctx.Query("email")
// 	name := ctx.Query("name")

// 	users, err := u.userUC.GetAllUser(ctx, page, size, email, name)
// 	if err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDTO.NewResponseSuccess(ctx, users, "success")
// }

// func (u *UserDelivery) getUserById(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	users, err := u.userUC.GetUserByID(ctx, id)
// 	if err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	} else {
// 		jsonDTO.NewResponseSuccess(ctx, users, "success")
// 	}
// }

// func (u *UserDelivery) updateUser(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	var userUpdates map[string]interface{}
// 	if err := ctx.ShouldBindJSON(&userUpdates); err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	err := u.userUC.UpdateUser(ctx, id, userUpdates)
// 	if err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDTO.NewResponseSuccess(ctx, nil, "success")
// }

// func (u *UserDelivery) deleteUser(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	err := u.userUC.DeleteUser(ctx, id)
// 	if err != nil {
// 		jsonDTO.NewResponseError(ctx, err.Error())
// 		return
// 	}

// 	jsonDTO.NewResponseSuccess(ctx, nil, "success")
// }
