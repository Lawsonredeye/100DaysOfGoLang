package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Person struct {
	Name   string `json:"name"`
	Age    int	`json:"age"`
	Emails []string `json:"emails"`
	Address string
}

func main() {
	json_bytes := []byte(`
		{
			"Name":"Lawson Redeye",
			"Age":27,
			"Emails":["lawson@mail.com","redeye@mail.com"],
			"Score":97
		}
	`)
	
	var pData map[string]interface{}
	err := json.Unmarshal(json_bytes, &pData)
	check(err)
	
	var emails []interface{}
	var firstEmail string
	emails = pData["Emails"].([]interface{})
	firstEmail = emails[0].(string)

	log.Println(strings.ToUpper(firstEmail))
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
