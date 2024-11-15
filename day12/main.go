package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello world\n")
	})

	http.HandleFunc("/shops", func(w http.ResponseWriter, r *http.Request){
		d := Data{"Boots", 1, "Best hiking boots & best selling in finland!"}
		res, err := json.Marshal(d)
		if err != nil {
			fmt.Println("Improper type")
		}
		w.Write(res)
	})

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server", err)
	}
}

type Data struct {
	Item string
	ItemId int
	ItemDesc string
}
