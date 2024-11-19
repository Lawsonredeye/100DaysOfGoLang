package main

import (
	"fmt"
	"net/http"
)

func aboutPage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/about.html")
}

func main() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/about", aboutPage)
	fmt.Println("Server is running at port 8080")
	http.ListenAndServe(":8080", nil)
}
