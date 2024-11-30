package main

import (
	"demoserver/myserver"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", myserver.Login)
	http.HandleFunc("/welcome", myserver.Welcome)

	log.Printf("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}