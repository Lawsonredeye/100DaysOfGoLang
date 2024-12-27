package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	r "todoapp/app"	
)

func main() {
	fmt.Println("Helloo todo app")
	router := gin.Default()

	router.POST("/task", r.CreateTask)
	router.Run()
}
