package main

import "log"

func main() {
	log.Println("Hello, Logger!")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Logger with microseconds")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")
}
