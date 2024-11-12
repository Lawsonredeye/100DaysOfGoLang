package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)


func main() {
		formData, err := json.Marshal(map[string]string{
		"name": "lawson",
		"address": "332 Main Avenue, Nigeria", // Come visit ;)
	})
	checkErr(err)
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	url := "https://httpbin.org/post"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(formData))
	request.Header.Set("Content-Type", "application/json")
	checkErr(err)
	
	resp, err := client.Do(request)
	checkErr(err)
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	checkErr(err)

	log.Println(data["data"].(string))
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
