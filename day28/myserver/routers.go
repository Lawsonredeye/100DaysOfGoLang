package myserver

import (
	model "demoserver/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Db connection variables
var URL = os.Getenv("db_user") + ":" + os.Getenv("db_pwd") + "@tcp(127.0.0.1)/moshdb"
var db = model.Conn(URL)

// Handles the authentication of users input passed from the request header
// and validates users credential.
func Login(w http.ResponseWriter, r *http.Request) {

	var (
		user     model.User
		validate model.User
	)

	defer r.Body.Close()
	res, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(res, &validate)
	if err != nil {
		log.Printf("error: %q", err)
		return
	}

	db.Table("users").Select("username", "password").Where("username = ?",
		validate.Username).Scan(&user)

	fmt.Printf("%+v\n", validate)
	fmt.Printf("%+v\n", user)

	if validate.Password != user.Password {
		http.Error(w, "401, user is not authorized", http.StatusUnauthorized)
		return
	}

	log.Printf("user %s, logged in successfully...", validate.Username)
	w.WriteHeader(200)
}

// Hello world for server.
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, user!")
}

func checkErr(e error, w http.ResponseWriter) {
	if e != nil {
		// log.Printf("%e", e)
		// }
		http.Error(w, "Bad request from user", http.StatusBadRequest)
		return
	}
}

// Handles User registrations.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var CreateUser model.User
	requestParam, err := io.ReadAll(r.Body)
	err = json.Unmarshal(requestParam, &CreateUser)
	checkErr(err, w)

	dbResult := db.Create(&CreateUser)
	checkErr(dbResult.Error, w)

	log.Printf("%s created an account.", CreateUser.Username)
	w.WriteHeader(201) // StatusCode indicating user account creation
}
