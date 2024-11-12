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
Here is a simple way of sending a post request to a server and seeing if the headers was well set:
```Go
func main() {
	packetData, err := json.Marshal(map[string]string{
		"name": "Lawson",
		"email": "gopheriscool@mail.com",
	})
	checkErr(err)
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewReader(packetData))
	checkErr(err)
	log.Println(resp.Header.Get("Content-Type"))

}
```

## POSTFORM
With Golang you can also PostForm data which would be needed in forms with a default content-type of application/x-www-form-urlencoded.
Here is a Demo:
```Go
func main() {
	// url.Values => map[string][]string
	formData := url.Values{
		"name": {"lawson"},
		"address": {"332 Main Avenue, Nigeria"}, // Come visit ;)
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	checkErr(err)
	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}
```

It is important to note that whilst http.Get allow you to create an easy way of sending request, it does not give you the flexibility to tweak your GET/POST requests. That is where the HTTP.CLIENT shines.

## USING http.Client{}
Here is a way to create your very own custom request with additional headers
```Go
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
```

## Resources
1. [Making Http requests in Golang](https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7)
2. [Sending URL values with http.PostForm in golang](https://pkg.go.dev/net/url#Values)
