package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func GenerateUUID() string {
	return uuid.New().String()
}

func Send200Response(ctx *gin.Context, message string, data interface{}) {
	res := Response{
		Message: message,
		Status:  "Success",
		Error:   "",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, res)
}

func Send201Response(ctx *gin.Context, message string, data interface{}) {
	res := Response{
		Message: message,
		Status:  "Success",
		Error:   "",
		Data:    data,
	}
	ctx.JSON(http.StatusCreated, res)
}

func Send400Response(ctx *gin.Context, message string, err string) {
	res := Response{
		Message: message,
		Status:  "Error",
		Error:   err,
		Data:    nil,
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func Send401Response(ctx *gin.Context, message string, err string) {
	res := Response{
		Message: message,
		Status:  "Error",
		Error:   err,
		Data:    nil,
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

func Send500Response(ctx *gin.Context, message string, err string) {
	res := Response{
		Message: message,
		Status:  "Error",
		Error:   err,
		Data:    nil,
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
}
