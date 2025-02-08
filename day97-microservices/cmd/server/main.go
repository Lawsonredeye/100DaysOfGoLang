package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Books struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Price float32 `json:"price"`
	PublishDate string `json:"publish_date"`
}

type Message struct {
	Greeting string `json:"greeting"`
}

var BookCatalog = []Books{}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/books", fetchAllBooks)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)// Allowing CORS
	message := Message{Greeting: "Hello, World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func fetchAllBooks(w http.ResponseWriter, r *http.Request){
	enableCors(w)// Allowing CORS
	if r.Method == "POST" {
		var newBook Books
		// defer r.Body.Close()
		json.NewDecoder(r.Body).Decode(&newBook)

		BookCatalog = append(BookCatalog, newBook)
		w.WriteHeader(http.StatusCreated)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BookCatalog)
}