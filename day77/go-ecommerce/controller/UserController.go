package controller

import "github.com/gin-gonic/gin"

func CreateAccount(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "account creation is successful",
	})
}
