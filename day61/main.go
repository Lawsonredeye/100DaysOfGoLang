package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(AuthBasicMiddleWare)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/admin", RBACMiddleware("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "welcome admin :)",
		})
	})
	r.Run()
}

func AuthBasicMiddleWare(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "valid-API-key" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unathorized request"})
		c.Abort()
		return
	}

	c.Next()
}

func BasicMiddleWare(c *gin.Context) {
	fmt.Println("I was passed and called!")
	c.Next()
}
