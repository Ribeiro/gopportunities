package handler

import (
	"fmt"
	"net/http"

	"github.com/Ribeiro/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, httpStatus int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(httpStatus, gin.H{
		"message": msg,
		"status":  httpStatus,
		"data":    nil,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"status":  http.StatusOK,
		"data":    data,
	})
}

type ErrorResponse struct {
	Message    string                  `json:"message"`
	Data       schemas.OpeningResponse `json:"data"`
	HttpStatus string                  `json:"status"`
}

type CreateOpeningResponse struct {
	Message    string                  `json:"message"`
	Data       schemas.OpeningResponse `json:"data"`
	HttpStatus string                  `json:"status"`
}

type DeleteOpeningResponse struct {
	Message    string                  `json:"message"`
	Data       schemas.OpeningResponse `json:"data"`
	HttpStatus string                  `json:"status"`
}
type ShowOpeningResponse struct {
	Message    string                  `json:"message"`
	Data       schemas.OpeningResponse `json:"data"`
	HttpStatus string                  `json:"status"`
}
type ListOpeningsResponse struct {
	Message    string                    `json:"message"`
	Data       []schemas.OpeningResponse `json:"data"`
	HttpStatus string                    `json:"status"`
}
type UpdateOpeningResponse struct {
	Message    string                  `json:"message"`
	Data       schemas.OpeningResponse `json:"data"`
	HttpStatus string                  `json:"status"`
}
