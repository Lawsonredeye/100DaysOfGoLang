package main

import (
	"go-commerce/controller"
	"go-commerce/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	model.DB.AutoMigrate(&model.Users{})

	r.POST("/signup", controller.CreateAccount)
	r.Run(":5050")
}
