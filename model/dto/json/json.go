package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type (
	jsonResponse struct {
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
	}

	jsonErrorResponse struct {
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

	paginationResponse struct {
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
		Message   string            `json:"responseMessage"`
		ErrorDesc []ValidationField `json:"error_description,omitempty"`
	}
)

func NewResponseSuccess(c *gin.Context, result interface{}, message string) {
	c.JSON(http.StatusOK, jsonResponse{
		Message: message,
		Data:    result,
	})
}

func NewResponseWithPaging(c *gin.Context, result interface{}, page int, total int) {
	c.JSON(http.StatusOK, paginationResponse{
		Data: result,
		Paging: pageData{
			Page:      page,
			TotalData: total,
		},
	})
}

func NewResponseValidationError(c *gin.Context, validationField []ValidationField, message string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Message:   message,
		ErrorDesc: validationField,
	})
}

func NewResponseBadRequest(c *gin.Context, err, message string) {
	c.JSON(http.StatusBadRequest, jsonErrorResponse{
		Message: message,
		Error:   err,
	})
}

func NewResponseError(c *gin.Context, err string) {
	errResp := errors.New(err)
	log.Error().Err(errResp).Msg(err)

	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Message: "internal server error",
		Error:   err,
	})
}

func NewResponseForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Message: message,
	})
}
