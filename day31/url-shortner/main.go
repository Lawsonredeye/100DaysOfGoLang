package main

import (
	"log"
	"net/http"
	"url-shortener/view"
)

func main() {
	http.HandleFunc("/create", view.CreateLink)
	log.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
