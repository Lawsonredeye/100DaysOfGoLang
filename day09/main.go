package main

import (
	"io"
	"log"
	"net/http"
)


func main() {
	resp, err := http.Get("https://httpbin.org/get")
	checkErr(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	log.Println(string(body))
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
