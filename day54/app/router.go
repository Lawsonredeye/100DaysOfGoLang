package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context){
	newTask := Task{}
	err := c.BindJSON(&newTask)

	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	fmt.Printf("%+v\n\n", newTask)
	c.Status(201)

}
