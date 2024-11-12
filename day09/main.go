package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)


func main() {
	// url.Values => map[string][]string
	formData := url.Values{
		"name": {"lawson"},
		"address": {"332 Main Avenue, Nigeria"}, // Come visit ;)
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	checkErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	checkErr(err)
	log.Println(data["form"])
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
