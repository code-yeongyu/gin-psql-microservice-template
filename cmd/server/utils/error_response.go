package utils

import (
	"gin_psql_microservice_template/cmd/server/errors"
	"gin_psql_microservice_template/cmd/server/forms"

	"github.com/gin-gonic/gin"
)

// AbortWithErrorResponse aborts the request with the given error
func AbortWithErrorResponse(c *gin.Context, statusCode int, errorType string, detail string) {
	errorResponse := forms.ErrorResponse{ErrorType: errorType, Message: errors.Messages[errorType], Detail: detail}
	c.AbortWithStatusJSON(statusCode, errorResponse)
}
