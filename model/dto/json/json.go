package jsonDTO

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
	jsonResponse struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
	}

	jsonResponseToken struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data,omitempty"`
	}

	jsonErrorResponse struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

	paginationResponse struct {
		Code   string      `json:"responseCode"`
		Data   interface{} `json:"data,omitempty"`
		Paging pageData    `json:"paging"`
	}

	pageData struct {
		Page      int `json:"page"`
		TotalData int `json:"totalData"`
	}
	ValidationField struct {
		FieldName string `json:"field"`
		Message   string `json:"message"`
	}
	jsonBadRequestResponse struct {
		Code      string            `json:"responseCode"`
		Message   string            `json:"responseMessage"`
		ErrorDesc []ValidationField `json:"error_description,omitempty"`
	}

	response struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
	}
)

func NewResponseSuccess(c *gin.Context, result interface{}, message string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    "200",
		Message: message,
		Data:    result,
	})
}

func NewResponseUserPaging(c *gin.Context, result interface{}, page int, total int) {
	c.JSON(http.StatusOK, paginationResponse{
		Code: "200",
		Data: result,
		Paging: pageData{
			Page:      page,
			TotalData: total,
		},
	})
}

func NewResponseSuccessToken(c *gin.Context, token, message string) {
	c.JSON(http.StatusOK, jsonResponseToken{
		Code:    "200",
		Message: message,
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	})
}

func NewResponseBadRequest(c *gin.Context, validationField []ValidationField, message string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Code:      "400",
		Message:   message,
		ErrorDesc: validationField,
	})
}

func NewResponseError(c *gin.Context, err string) {
	log.Error().Msg(err)

	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Code:    "500",
		Message: "internal server error",
		Error:   err,
	})
}

func NewResponseForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403",
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401",
		Message: message,
	})
}

func NewResponseAuth(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, response{
		Code:    "401",
		Message: message,
	})
}
