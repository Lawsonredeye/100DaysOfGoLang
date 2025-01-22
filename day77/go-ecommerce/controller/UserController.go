package controller

import (
	"fmt"
	"go-commerce/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

// Login handles the log in of users for having access into db.
// Accepts POST method with a ContentType: application/json with body containing details
// e.g username or email, and password.
// Stores a cookie session on the header where JWT would be stored.
// Parameter:
// - c : *gin.Context - GIN framework
// Response:
// - HTTP 200: Successfully signed in. 
// - HTTP 401: User is unauthorized.
// - HTTP 500: Unable to assign JWT
func Login(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Content-Type Header",
		})
		return
	}

	var authUser model.Users
	var foundUser model.Users
	c.BindJSON(&authUser)

	if authUser.Username == "" || authUser.PasswordHash == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"Username or password field cant be empty"})
		return
	}

	// check if user exists in the db, if not return user not found error
	model.DB.Where("username = ?", authUser.Username).First(&foundUser)

	// if result.Error != nil {
	// 	c.AbortWithStatus(400)
	// 	return
	// }

	// check if the password is valid if not return 401 unauthorized
	if bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(authUser.PasswordHash)) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"User is not authorized"})
		return
	}

	// if user is authorized then create a jwt token and store in the user header
	token, err := createToken(authUser.Username)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to assign token for user"})
		return
	}
	c.SetCookie(authUser.Username, token, 24000, "", "", false, true)

	// return response that user is logged in successfully.
	c.JSON(http.StatusCreated, gin.H{"msg":"Login successful."})
	
}

// createToken creates a signed JWT token for granting logged in users access.
// to perfrom certain functions
// Parameter:
// 	username - string
// Response:
// - string: signed token string
// - error : nil or non-nil value
func createToken(username string) (string, error) {
	err := godotenv.Load()

	key := os.Getenv("SECRET_KEY")
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 576).Unix(),
	})

	// convert token to byte String when signing
	tokenString, err := token.SignedString([]byte(key))

	fmt.Println(token)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err:= jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
		return os.Getenv("SECRET_KEY"), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	
	return nil
}
