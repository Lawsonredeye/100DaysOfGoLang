package main

import (
	"io"
	"log"
	"net/http"
)

const URL string = "https://localhost:4000/protected"

func main() {
	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)

	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("Lawsonredeye", "gogophers")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("%s", string(body))
}
