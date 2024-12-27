package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	newTask := Task{}
	err := c.BindJSON(&newTask)

	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	newTask.ID = len(AllTasks)

	fmt.Printf("%+v task added successfully!\n\n", newTask)

	AllTasks = append(AllTasks, newTask)
	c.Status(201)
}

func GetTask(c *gin.Context) {
	c.JSON(200, AllTasks)
	log.Printf("%+v\n", AllTasks)
}
