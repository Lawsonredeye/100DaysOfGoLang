package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func main() {
	postBody, _ := json.Marshal(map[string]string{
		"name": "Lawson",
		"email": "lawson@mail.com",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.NewRequest(http.MethodPost, "https://postman-echo.com/post", responseBody)
	resp.Header.Add("X-Auth-Key", "HERE_IS_KEY")

	response, err := http.DefaultClient.Do(resp)

	if err != nil {
		log.Fatalf("An error occured %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
