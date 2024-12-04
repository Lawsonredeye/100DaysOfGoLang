package main

import (
	"log"
	"net/http"
	"url-shortener/view"
)

func main() {
	http.HandleFunc("/create", view.CreateLink)
	http.HandleFunc("/", view.RedirectLink)
	log.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
