package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// create the url
	url := "https://jsonplaceholder.typicode.com/posts/1"
	request, err := http.NewRequest("GET", url, nil)
	CheckError(err)
	// create the client which would send the request
	client := &http.Client{}
	// send the get request
	response, err := client.Do(request)
	CheckError(err)
	// get the data and convert from a byte into a byte slice
	resBody, err := io.ReadAll(response.Body)
	CheckError(err)
	// printout the data from from the response
	var b bytes.Buffer
	err = json.Indent(&b, resBody, "", "  ")
	CheckError(err)
	fmt.Printf("%+v\n", string(b.Bytes()))
	// defer response.Body.Close()
}
