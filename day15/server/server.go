package main

import (
	"fmt"
	"net/http"
	"strings"
)

func welcomeMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func searchFunc(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if q.Get("q") == "" {
		http.Error(w, "Query parameter is missing", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Searching for: %s", q.Get("q"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Path

	pathLen := strings.TrimPrefix(r.URL.Path, "/user/") 
	if len(pathLen) < 3  {
		http.Error(w, "Invalid User ID.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "User ID: %v", pathLen)
}

func main() {
	http.HandleFunc("/{$}", welcomeMsg)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/search", searchFunc)
	http.HandleFunc("/user/", userHandler)

	fmt.Println("Serving content at port 8080")
	http.ListenAndServe(":8080", nil)
}

