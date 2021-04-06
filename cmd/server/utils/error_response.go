package utils

import (
	"gin_psql_microservice_template/cmd/server/errors"
	"gin_psql_microservice_template/cmd/server/forms"

	"github.com/gin-gonic/gin"
)

// AbortWithErrorResponse aborts the request with the given error
func AbortWithErrorResponse(c *gin.Context, statusCode int, errorType errors.CustomError, detail string) {
	errorResponse := forms.ErrorResponse{ErrorType: string(errorType), Message: errors.Messages[errorType], Detail: detail}
	c.AbortWithStatusJSON(statusCode, errorResponse)
}
