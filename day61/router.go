package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetHeader("Role")

		if userRole != requiredRole {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized to carry out operations on this route!",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
