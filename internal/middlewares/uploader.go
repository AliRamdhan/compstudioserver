package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateFileMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if a file is uploaded
		fileHeader, err := c.FormFile("ProductImage")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			c.Abort()
			return
		}
		// Set the file header in the context
		c.Set("fileHeader", fileHeader)
		c.Next()
	}
}
