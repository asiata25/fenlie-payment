package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
	// JSONResponse - struct for JSON response success
	jsonResponse struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
	}

	// JSONResponse - struct for JSON response success token
	jsonResponseToken struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data,omitempty"`
	}

	// JSONResponse - struct for JSON response error
	jsonErrorResponse struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

	ValidationField struct {
		FieldName string `json:"field"`
		Message   string `json:"message"`
	}

	// JSONResponse - struct for JSON response bad request
	jsonBadRequestResponse struct {
		Code             string            `json:"responseCode"`
		Message          string            `json:"responseMessage"`
		ErrorDescription []ValidationField `json:"error_description,omitempty"`
	}

	// JSONResponse - struct for JSON response auth
	Response struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
	}
)

func NewResponseSuccess(c *gin.Context, result interface{}, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    "200" + serviceCode + responseCode,
		Message: message,
		Data:    result,
	})
}

func NewResponseSuccessToken(c *gin.Context, token, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponseToken{
		Code:    "200" + serviceCode + responseCode,
		Message: message,
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	})
}

func NewResponseBadRequest(c *gin.Context, validationField []ValidationField, message, serviceCode, errorCode string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Code:             "400" + serviceCode + errorCode,
		Message:          message,
		ErrorDescription: validationField,
	})
}

func NewResponseError(c *gin.Context, err, serviceCode, errorCode string) {
	log.Error().Msg(err)
	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Code:    "500" + serviceCode + errorCode,
		Message: "Internal Server Error",
		Error:   err,
	})
}

func NewResponseForbidden(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseAuth(c *gin.Context, message, serviceCode, errorCode string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Code:    "401" + serviceCode + errorCode,
		Message: message,
	})
}
