package controller

import (
	"go-commerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	// check if the incoming request is a JSON application
	if c.ContentType() != "application/json" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// create a user object for handling the json data
	var user model.Users

	c.BindJSON(&user)

	if user.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Empty username field."})
		return
	}
	if user.PasswordHash == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Empty password field."})
		return
	}

	if user.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Empty email field."})
		return
	}

	// check if the following fields are empty ['username', 'email', 'password'] && return 400 if empty

	// check if the username already exists in the db && return 401 if exists
	model.DB.Where()
	// hash password using bycrypt and then write to the go user object

	// using the gorm connection, save the new instance of the data and update the userId

	c.JSON(201, gin.H{
		"message": "account creation is successful",
	})
}
