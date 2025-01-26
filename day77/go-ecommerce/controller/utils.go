package controller

import (
	"go-commerce/model"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var CategoryAbbreviation = map[string]string{
	"jeans":  "JNS",
	"tshirt": "TSH",
	"jacket": "JKT",
	"shoes":  "SHS",
	"shorts": "SRT",
	"hoodie": "HDY",
	"dress":  "DRS",
	"skirt":  "SKT",
	"suit":   "SUT",
	"hat":    "HAT",
	"laptop": "LPT",
}

var ColorCode = map[string]string{
	"white":   "WHT",
	"red":     "RED",
	"green":   "GRN",
	"blue":    "BLU",
	"yellow":  "YLW",
	"pink":    "PNK",
	"brown":   "BRW",
	"magenta": "MGT",
	"rose":    "RSE",
	"velvet":  "VEL",
	"black":   "BLK",
}

type Obj struct {
	ProductID int `json:"product_id"`
	Quantity int `json:"quantity"`
}

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}


func ValidateCookies(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	var user model.Users
	model.DB.Where("username = ?", token.Claims.(*MyCustomClaims).Username).First(&user)
	return int(user.ID),nil
}
