package main

import (
	"go-commerce/controller"
	"go-commerce/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	model.DB.AutoMigrate(&model.Users{})

	r.POST("/api/v1/signup", controller.CreateAccount)
	r.POST("/api/v1/login", controller.Login)
	r.GET("/api/v1/logout", controller.Logout)
	r.POST("/api/v1/product", controller.AddProduct)
	r.GET("/api/v1/product/:name", controller.FindProductByName)
	r.DELETE("/api/v1/product/:id", controller.DeleteProductByID)
	r.PUT("/api/v1/product/:id", controller.UpdateProductByID)
	r.POST("/api/v1/orders", controller.AddOrder)
	r.Run(":5050")
}
