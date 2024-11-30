package myserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name     string
	Password string
}

type DBUser struct {
	Name     string
	Birthday time.Time
	Age      int
	ID       uint
}

var ExistingUser = User{
	Name:     "lawson",
	Password: "DareDevilGopher",
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var validate User
	res, _ := io.ReadAll(r.Body)
	// user:pwd@tcp(localhost:3306)/moshdb
	url := os.Getenv("db_user") + ":" + os.Getenv("db_pwd") + "@tcp(localhost:3306)/moshdb"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	checkErr(err)

	var user DBUser

	db.Model(DBUser).Find(&user)

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

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
