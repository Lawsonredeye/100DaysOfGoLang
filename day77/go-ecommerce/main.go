package main

import (
	"go-commerce/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/signup", controller.CreateAccount)
	r.Run(":5050")
}
