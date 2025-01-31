package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Hello world %v", counter)
	counter++
	mutex.Unlock()
}


func main() {
	http.HandleFunc("/hi", incrementCounter)
	http.HandleFunc("/", echoString)
	 log.Fatal(http.ListenAndServe(":8080", nil))
}
