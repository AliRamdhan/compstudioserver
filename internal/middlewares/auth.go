package middlewares

import (
	"github.com/AliRamdhan/compstudioserver/auth"
	"github.com/gin-gonic/gin"
)

// AdminAuth middleware for admin access
func AdminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		// Validate the token
		claims, err := auth.ParseTokenClaims(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		// Check if the user has admin role
		if claims.RoleID != 1 {
			context.JSON(403, gin.H{"error": "forbidden: insufficient permissions"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// UserAuth middleware for user access
func UserAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		// Validate the token
		claims, err := auth.ParseTokenClaims(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		// Check if the user has user role
		if claims.RoleID != 2 {
			context.JSON(403, gin.H{"error": "forbidden: insufficient permissions"})
			context.Abort()
			return
		}
		context.Next()
	}
}
