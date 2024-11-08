# Struct and Interfaces - Day 05

## Struct
This is used to create user defined data types which is similar to classes in other languages.
To create a struct you use the `type name_of_struct struct {...}` syntax, where in between the struct you define members of the struct.

Here is a demo for creating structs:
```Go
type MyFifaTeam struct {
    ClubName string
    ClubCaptain string
    ClubCoach string
    ClubTeamMembers int
    ...
}
```
Struct members can be accessed using the dot(.) notation similar to the ones in python classes.
Here is how to access members of a struct:
```Go
var manCity MyFifaTeam
manCity.ClubName = "Manchester City"
manCity.ClubCoach = "Pep Guardiola"
...
```
Struct also have struct tags which are use to define a struct when working with external libraries like the encoding/json library.


# Interface
An interface is used to create methods for objects with similar properties.
To create an Interface you use the `type interface_name interface {...}`. In between the curly brackets is where you define the methods that objects of an interface needs to meet.

Here is an example:
```Go
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
```

## Resources
1. [Interfaces](https://www.simplilearn.com/tutorials/golang-tutorial/guide-to-golang-interface)
2. [Interfaces explained](https://www.alexedwards.net/blog/interfaces-explained)
