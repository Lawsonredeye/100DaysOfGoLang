package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/hello.html", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, r.URL.Path[1:])	
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Can't run on port 8080")
	} else {
		log.Println("running on port 8080")
	}
}
