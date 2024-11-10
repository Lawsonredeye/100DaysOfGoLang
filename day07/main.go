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
	// url := "https://jsonplaceholder.typicode.com/posts/1"
	// request, err := http.NewRequest("GET", url, nil)
	// CheckError(err)
	// // create the client which would send the request
	// client := &http.Client{}
	// // send the get request
	// response, err := client.Do(request)
	// fmt.Println(response.Status)
	// CheckError(err)
	// // get the data and convert from a byte into a byte slice
	// resBody, err := io.ReadAll(response.Body)
	// CheckError(err)
	// // printout the data from from the response
	// var b bytes.Buffer
	// err = json.Indent(&b, resBody, "", "  ")
	// CheckError(err)
	// fmt.Printf("%+v\n", string(b.Bytes()))
	// defer response.Body.Close()

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	CheckError(err)
	body, err := io.ReadAll(res.Body)
	CheckError(err)
	sb := string(body)
	fmt.Printf(sb)
	fmt.Println()
	defer res.Body.Close()

	// sending post request
	data, err := json.Marshal(map[string]string{
		"name": "Lawsonredeye",
		"email": "lawson@mail.com",
	})
	CheckError(err)
	dataBuffer := bytes.NewBuffer(data)
	// sending post data
	res, err = http.Post("https://postman-echo.com/post", "application/json",dataBuffer)
	CheckError(err)
	defer res.Body.Close()
}
