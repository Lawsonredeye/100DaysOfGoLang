package controller

import (
	"go-commerce/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateAccount handles the creation of a new user account.
// Use the POST request and accept a content-type: application/json,
// typically contain the user details e.g (username, password, email).
// This function validates the email and username
// Parameter:
// - c: *gin.Context : GIN framework context manager
// Response:
// - HTTP 201: Account created successfully.
// - HTTP 400: User Details already exists and user account was not created.
// - HTTP 501: User account creation failed due to cryptograph hashing
func CreateAccount(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Content-Type Header",
		})
		return
	}
	// create a user object for handling the json data
	var user model.Users
	c.BindJSON(&user)

	// check if the following fields are empty ['username', 'email', 'password']
	// && return 400 if empty
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

	// check if the username and email already exists in the db
	// && return 401 if exists
	var resultUser model.Users
	model.DB.Where("username = ?", user.Username).First(&resultUser)

	if resultUser.Username != "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Username already exists",
		})
		return
	}
	
	var emailUser model.Users
	model.DB.Where("email = ?", user.Email).First(&emailUser)

	if emailUser.Email != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Account with email already exists.",
		})
		return
	}
	// hash password using bycrypt and then write to the go user object
	pwd := user.PasswordHash
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	user.PasswordHash = string(newHashedPassword)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error": "Password hashing failed.",
		})
		return
	}
	// using the gorm connection, save the new instance of the data and update the userId
	result := model.DB.Create(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Account not created, try again.",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "account creation is successful",
	})
}
