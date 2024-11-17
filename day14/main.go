package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func serveContent(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func sayHi(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi, client! Welcome to my web server.")
}

func main() {
	 http.HandleFunc("/", serveContent)
	 http.HandleFunc("/hi", sayHi)
	 log.Fatal(http.ListenAndServe(":8080", nil))
}
