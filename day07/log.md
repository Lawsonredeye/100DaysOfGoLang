# Sending Simple Get Request Using net/http - Day 07

So one of the major thing to do in every language is knowing how to send http requests and in golang it is no different.

Instead of using external libraries for sending basic network request, GoLang Standard Library can send Http network request without you breaking a sweat.

To send a basic http request you need to import the net/http library and then create a mock request which would prepare the client to send the request.

Next would be to call the client and provide the basic headers needed to send the request. Here is a code sample to demonstrate it:

```Go
package main

import (
    "fmt"
    "net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	request, err := http.NewRequest("GET", url, nil)
	CheckError(err)
	// create the client which would send the request
	client := &http.Client{}
	// send the get request
	response, err := client.Do(request)
	CheckError(err)
    fmt.Println(response.Status)
    // Output: 200 OK
}
```
What the above does is to send a get request and then prints the status code of the response.

## Resource
1. [How to send http request in golang](https://www.makeuseof.com/go-make-http-requests/#:~:text=Sending%20Requests%20Using%20Common%20HTTP%20Methods&text=It%20is%20a%20sub%2Dpackage,Go%20net%2Fhttp%20package's%20http.)
2. [Test your api with the jsonplaceholder](https://jsonplaceholder.typicode.com/posts)
3. [net/http library](https://pkg.go.dev/net/http)
