package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type Comment struct {
	PostId int `json:"postId"`
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string	`json:"body"`
}

// Playing with encoding and decoding of json data using the Go Std library.
func main() {
	var data []Comment
	const URL string = "https://jsonplaceholder.typicode.com/comments/id/4" 

	resp, err := http.Get(URL)
	check(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	
	// log.Printf("body: %T", body)
	err = json.Unmarshal(body, &data)
	check(err)

	for idx := range data {
		log.Printf("%v", strings.ToUpper(data[idx].Email))
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
