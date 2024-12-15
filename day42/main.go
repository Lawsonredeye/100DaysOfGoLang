package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Hello, Logger!")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Logger with microseconds")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Setting up a new custom logger object in golang
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)

	// Also setting the flags for our new custom logger object using the SetFlags() method
	mylog.SetFlags(log.LstdFlags | log.Lmicroseconds)
	mylog.Println("using mylog logger")

	// Setting prefixs for custom logger or default logger
	mylog.SetPrefix("web-server-access.log: ")
	mylog.Println("Hello, from my custom logger")

	// custom outputs and logger objects of io.Writer
	var buf bytes.Buffer

	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("Hello from custom outputs")
	fmt.Print("from buflog:", buf.String())
}
