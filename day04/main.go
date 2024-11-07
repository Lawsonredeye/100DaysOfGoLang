package main

import (
	"fmt"
	"strings"
)

func main() {

	// creating an array
	seaCreatures := [3]string{"shark", "goldfish", "seahorse"}
	// accessing elements of an array
	fmt.Println(seaCreatures[2])

	// creating a slice
	landCreatures := []string{"Elephant", "Rhinos", "python", "gopher"}
	fmt.Println(landCreatures[3])

	// creating a username map
	usernames := map[string]string{
		"lawson": "redeye",
		"george": "escay01",
		"golang": "gopher",
		"python": "py",
		"kelvin": "red ninja",
	}
	user := ""
	
	fmt.Println("Enter a valid username:")
	if _, err := fmt.Scanln(&user); err != nil {
		fmt.Println("you didnt pass a username")
	}
	user = strings.TrimSpace(user)

	if u, ok := usernames[user]; !ok {
		fmt.Println("user not found")
	} else {
	fmt.Println("welcome,", u)
	}

	// delete from a map
	delete(usernames, "python")
	fmt.Println(usernames)
}

