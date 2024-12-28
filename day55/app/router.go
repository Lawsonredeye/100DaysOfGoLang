package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	stringID := c.Param("id")
	id, err := strconv.Atoi(stringID)

	if err != nil {
		c.AbortWithStatusJSON(401, "Invalid Task ID")
		return
	}

	for _, val := range AllTasks {
		if id == val.ID {
			c.JSON(200, val)
			return
		}
	}
	res := map[string]string{"error": "Task not found"}
	c.AbortWithStatusJSON(402, res)
}

// UpdateTaskByID updates stored tasks by its ID.
func UpdateTaskByID(c *gin.Context) {
	stringID := c.Param("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, val := range AllTasks {
		if id == val.ID {
			newData := Task{}
			err := c.BindJSON(&newData)
			if err != nil {
				c.AbortWithError(http.StatusGatewayTimeout, err)
				return
			}
			val.Title = newData.Title

			val.Description = newData.Description
			AllTasks[id] = val
			c.JSON(http.StatusOK, val)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotFound, map[string]string{"error": "Task with id not found"})
}

// DeleteTaskByID removes a specific task from the db in-memory storage.
func DeleteTaskByID(c *gin.Context) {
	paramID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newTaskA := AllTasks[:paramID]
	newTaskB := AllTasks[paramID+1:]
	AllTasks = newTaskA
	AllTasks = append(AllTasks, newTaskB...)
	c.Status(http.StatusOK)
}
