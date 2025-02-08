package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type Message struct {
	Greeting string `json:"greeting"`
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Greeting: "Hello, World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}