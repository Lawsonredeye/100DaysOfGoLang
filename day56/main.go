package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", GetUser)
	r.Run()
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":     "lawsonredeye",
		"codename": "redeye007",
	})
}
