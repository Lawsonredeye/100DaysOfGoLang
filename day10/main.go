package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name   string `json:"name"`
	Age    int	`json:"age"`
	Emails []string `json:"emails"`
}

func main() {
	emails := []string{"lawson@mail.com", "redeye@mail.com"}
	json_bytes, err := json.Marshal(emails)
	check(err)
	log.Printf("%s", json_bytes)
	// outputs:
	// 2024/11/13 13:31:36 ["lawson@mail.com","redeye@mail.com"]
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
