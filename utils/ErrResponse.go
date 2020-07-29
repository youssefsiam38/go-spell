package utils

import (
	"github.com/gin-gonic/gin"
)

// ErrResponse send a response with error message and status code
func ErrResponse(c *gin.Context, code int, message string) {
	c.JSON(code, struct{
		Message string `json:"message"`
	}{
		Message: message,
	})
}