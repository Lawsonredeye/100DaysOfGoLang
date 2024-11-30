package myserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name     string
	Password string
}

var ExistingUser = User{
	Name:     "lawson",
	Password: "DareDevilGopher",
}

func Login(w http.ResponseWriter, r *http.Request) {

	// r.Body.Close()

	// username := r.PostFormValue("name")
	// userpwd := r.PostFormValue("password")

	// fmt.Println(username, userpwd)
	// if username != User.Name && userpwd != User.Password {
	// 	http.Error(w, "401, User not authorized", http.StatusUnauthorized)
	// 	return
	// }

	// w.WriteHeader(200)

	defer r.Body.Close()
	var validate User
	res, _ := io.ReadAll(r.Body)

	_ = json.Unmarshal(res, &validate)
	// fmt.Println(validate)

	if validate.Name != ExistingUser.Name && validate.Password != ExistingUser.Password {
		http.Error(w, "401, user is not authorized", http.StatusUnauthorized)
		return
	}

	log.Printf("user %s, logged in successfully...", validate.Name)
	w.WriteHeader(200)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, user!")
}
