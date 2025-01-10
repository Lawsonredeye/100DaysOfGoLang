package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var JWT_KEY = []byte("Go go gophers")

var AllUsers = make([]LoginDetails, 0)

func Login(c *gin.Context) {
	data := LoginDetails{
		Username: "Lawsonredeye",
		Password: "ILOVEGOLANG",
	}

	AllUsers = append(AllUsers, data)
	user := LoginDetails{}
	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	foundUser, err := FindUser(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found."})
	}

	if user.Password != foundUser.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_KEY)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func FindUser(u LoginDetails) (LoginDetails, error) {

	for _, val := range AllUsers {
		if u.Username == val.Username {
			return val, nil
		}
	}

	return LoginDetails{}, nil
}
