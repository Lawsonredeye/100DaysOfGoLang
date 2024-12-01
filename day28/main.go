package main

import (
	"demoserver/myserver"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", myserver.RegisterUser)
	http.HandleFunc("/welcome", myserver.Welcome)
	http.HandleFunc("/login", myserver.Login)

	log.Printf("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
