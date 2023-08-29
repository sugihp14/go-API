package response

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message"`
}

func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func ErrorJSON(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Message: message})
}
