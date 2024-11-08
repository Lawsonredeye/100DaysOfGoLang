package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// learning struct and struct tags


type User struct {
	Name string `json:"name"`
	Password string `json:"-"`
	PrefferedFish []string `json:"preferredFish ,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name: "Sammy the shark",
		Password: "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
	
	// Learning Interfaces
	fmt.Println("Interface")

	// create a method that updates the number of books in a shelf
	// var porsche motorTuning

	porsche := Car{"gtr 911", 32.6, 44.2, 3.0}
	printSpec(porsche)
	porsche.Speed()
	porsche.Hp()
	printSpec(porsche)
}

type motorTuning interface {
	Speed() float64
	Hp() float64
}

type Car struct {
	Model      string
	HorsePower float64
	Acceleration float64
	Brake      float64
}

func (c Car) Hp() float64 {
	c.HorsePower *= 2.0
	return c.HorsePower
}

func (c Car) Speed() float64 {
	c.Acceleration += 10.9
	return c.Acceleration
}

func printSpec(m motorTuning) {
	fmt.Printf("Current Specs\n\tSpeed: %v\n\tHorsepower(hp): %v\n", m.Speed(), m.Hp())
}
