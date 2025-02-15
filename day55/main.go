package main

import (
	r "todoapp/app"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/tasks", r.CreateTask)
	router.GET("/tasks", r.GetTask)
	router.GET("/tasks/:id", r.GetTaskByID)
	router.PATCH("/tasks/:id", r.UpdateTaskByID)
	router.DELETE("/tasks/:id", r.DeleteTaskByID)
	router.Run()
}
