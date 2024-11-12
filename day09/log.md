# More Http Request Using NET/HTTP - Day 09

## GET
Here is a simple way of sending a get request using the net/http library

```Go
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
```
Note:
We then defer the execution of resp.Body.Close() which will be executed at the end of the function.

This step is very very important. When you have a response body (it is not nil), forgetting to close the response body can cause resource leaks in a long running programs.

## POST



## Resources
1. [Making Http requests in Golang](https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7)
