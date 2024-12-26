package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/:id", getParameter)

	router.Run()
}

func home(c *gin.Context){
	c.String(http.StatusOK, "Hello from Gin!")
}

func getParameter(c *gin.Context){
	id := c.Param("id")
	msg := fmt.Sprintf("Your inputted ID is : %v\n\n", id)
	c.String(http.StatusOK, msg)
}
