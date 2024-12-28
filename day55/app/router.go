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

// GetTask fetches all tasks from the in-memory database.
func GetTask(c *gin.Context) {
	c.JSON(200, AllTasks)
	log.Printf("%+v\n", AllTasks)
}

// GetTaskByID fetches the task by the specified ID from the in memory db (array).
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	for _, val := range AllTasks {
		if id == val.ID {
			c.JSON(200, val)
			return
		}
	}
	res := map[string]string{"error": "Task not found"}
	c.AbortWithStatusJSON(402, res)
}
