package main

import (
	"log"
	"net/http"
	"url-shortener/view"
)

func main() {
	http.HandleFunc("/create", view.CreateLink)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
