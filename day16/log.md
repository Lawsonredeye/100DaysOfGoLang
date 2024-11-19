# Practice, Practice, Practice

For today i didnt do other things or learn a new concept but worked on my skills and knowledge on how to use the net/http standard library.

I learned about the importance of using http.StripPrefix() which takes a file/folder that you want to strip off the prefix to prevent clients from accessing other files or folders which should not be accessed at all.

I also learned more on how to serve a directory using the http.FileServer() to serve content from a specific directory.

Here is a simple demo of how to serve files from a directory in golang:
```Go
package main

import (
	"fmt"
	"net/http"
)

func aboutPage(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/about.html")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/about", aboutPage)
	fmt.Println("Server is running at port 8080")
	http.ListenAndServe(":8080", nil)
}
```
With just that you can be able to serve files from a directory which is located in your current project.
You could also choose to respond with files from a different directory but you will have to specify the entire path.

## Assignment
Goal: Build a web server that serves:

1. A simple HTML file (index.html) styled with CSS and with a button linked to JavaScript for interactivity.
2. Add a /static/ route to serve the static directory.

**Requirements: **

1. Create an HTML page with at least:
A heading.
A styled button.
A JavaScript function that interacts with the user.
2. Use http.FileServer for serving files.
3. Use http.ServeFile to explicitly serve index.html.

## Solution
[File that serves a web server request for a specific endpoint](https://github.com/Lawsonredeye/100DaysOfGoLang/blob/main/day16/static_server/main.go)
