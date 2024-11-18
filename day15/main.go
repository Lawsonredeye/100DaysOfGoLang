package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stats struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func welcomeMsg(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path == "/" {
	// fmt.Fprintf(w, "Hello, Lawson!")
	// } else {
	// http.Error(w, "status not found", http.StatusNotFound)
	// return
	// }

	fmt.Fprintf(w, "Hello, lawson!")	
}

func status(w http.ResponseWriter, r *http.Request) {
	// should return the status and uptime of the server
	res := Stats{"Server is running", "1 hour"}
	data, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
query := r.URL.Query().Get("q")
	if query == "" {
		fmt.Fprintf(w, "Please provide a search query!")
		return
	}
	fmt.Fprintf(w, "You searched for: %s", query)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.RawPath
	fmt.Fprintf(w, "User ID: %s", path)
	fmt.Println(r.URL.RawPath)
}

func main() {
	http.HandleFunc("/{$}", welcomeMsg)
	http.HandleFunc("/status", status)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/search", searchHandler)

	fmt.Println("Service is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
