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
	// fmt.Fprintf(w, "%v", string(data))
	// io.WriteString(w, data)
}

func main() {
	http.HandleFunc("/{$}", welcomeMsg)
	http.HandleFunc("/status", status)

	fmt.Println("Service is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
