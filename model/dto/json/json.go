<<<<<<< HEAD
package jsonDTO
=======
package json
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
<<<<<<< HEAD
=======
	// JSONResponse - struct for JSON response success
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
	jsonResponse struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
	}

<<<<<<< HEAD
=======
	// JSONResponse - struct for JSON response success token
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
	jsonResponseToken struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data,omitempty"`
	}

<<<<<<< HEAD
=======
	// JSONResponse - struct for JSON response error
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
	jsonErrorResponse struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

<<<<<<< HEAD
=======
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

>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
	paginationResponse struct {
		Code   string      `json:"responseCode"`
		Data   interface{} `json:"data,omitempty"`
		Paging pageData    `json:"paging"`
	}

	pageData struct {
		Page      int `json:"page"`
		TotalData int `json:"totalData"`
	}
<<<<<<< HEAD
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
=======
)

func NewResponseSuccess(c *gin.Context, result interface{}, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    "200" + serviceCode + responseCode,
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Message: message,
		Data:    result,
	})
}

<<<<<<< HEAD
func NewResponseUserPaging(c *gin.Context, result interface{}, page int, total int) {
	c.JSON(http.StatusOK, paginationResponse{
		Code: "200",
=======
func NewResponseUserPaging(c *gin.Context, result interface{}, page int, total int, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, paginationResponse{
		Code: "200" + serviceCode + responseCode,
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Data: result,
		Paging: pageData{
			Page:      page,
			TotalData: total,
		},
	})
}

<<<<<<< HEAD
func NewResponseSuccessToken(c *gin.Context, token, message string) {
	c.JSON(http.StatusOK, jsonResponseToken{
		Code:    "200",
=======
func NewResponseSuccessToken(c *gin.Context, token, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponseToken{
		Code:    "200" + serviceCode + responseCode,
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Message: message,
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	})
}

<<<<<<< HEAD
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
=======
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
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Error:   err,
	})
}

<<<<<<< HEAD
func NewResponseForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403",
=======
func NewResponseForbidden(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403" + serviceCode + errorCode,
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Message: message,
	})
}

<<<<<<< HEAD
func NewResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401",
=======
func NewResponseUnauthorized(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401" + serviceCode + errorCode,
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Message: message,
	})
}

<<<<<<< HEAD
func NewResponseAuth(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, response{
		Code:    "401",
=======
func NewResponseAuth(c *gin.Context, message, serviceCode, errorCode string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Code:    "401" + serviceCode + errorCode,
>>>>>>> fb16f33 (feat: category endpoint for get, get by id, create, delete, and update)
		Message: message,
	})
}
