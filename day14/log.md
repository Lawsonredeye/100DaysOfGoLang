# More on web servers
Web servers are actually one of the strong suite of golang but learning how to create a web server and understanding how a web server should behave is key.

For today, i'll be doing a little on how to create web servers that serves content, different types of contents using the nt/http library.

Here is a basic http web server that increments a counter when it is called while running:
```Go
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Hello world %v", counter)
	counter++
	mutex.Unlock()
}


func main() {
	http.HandleFunc("/hi", incrementCounter)
	http.HandleFunc("/", echoString)
	 log.Fatal(http.ListenAndServe(":8080", nil))
}
```

