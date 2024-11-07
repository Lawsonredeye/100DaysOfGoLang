# Advanced Go Data types - Day 04

## Arrays and slices
Arrays is a data structure that is used for storing values of same type. Unlike arrays from other programming language, the array in GoLang must be declared with its size to inform the compiler before runtime the amount of memory the declared array has.

This is majorly a constraint when it comes to using an array for storing large elements and wanting to expand the size further.

Arrays are declared in go in this format:
```Go
// var [name of array] [size]type
var arr [4]int
```
The above declares an array that stores only integers and by default the values of the arrays are zeros (0). Adding a non integer data type would cause a panic. 
The array also has a function `len()` which is used to check for the size of an array.

### Slices
Slices is exactly like an array but the only difference is that a slice doesn't have a fixed size, if declared with a fixed size it would then become an arry.

But that doesn't mean that a slice can't have a size, here is how to create a slice with a size using the `make()` function.
```Go
slice := make([]int, 4) // the second parameter acts as the length and if a third parameter is added it turns into the capacity of the slice
```
The above creates a slice with a length of 4 but don't be confused that it would become an array. As a matter of fact, the size of the slice can still increase and decrease if more elements are added.

An array can turn into a slice by performing a slicing technique which is seen in other language like py.
```Go
arr := [4]int{1, 2, 3, 4}
slice := make([]int, 4)

slice = arr[:]
```
The `arr[:]` means selecting the element from the range of the first element to the last element of the array. When such action is performed it turns the array into a slice and remember that in go a variable of a specific data type cant be change after its declared and thats why the above example was stored in another variable (a slice).

The slice has an `append()` which is used for joining an array or a slice to another slice. The append() doesnt work on an array, it must be a slice.
```Go
package main

import "fmt"

func main() {
	fruits := [4]string{"mango", "peach", "grape", "banana"}
	basket := []string{"papaya"}
	basket = append(basket, fruits[:]...) // using fruits[:]... spreads the fruits elements into the append func
	fmt.Println(basket)
    // output: [papaya mango peach grape banana]
}
```

The len() function can be used on the slice but it only returns the len of the slice which was defined when using make but not the capacity. To get the real size of a slice use the `cap()` function which returns the total capacity of memory the slice has.

Also two dimensional slice/array can be created in a simple way in Go. Here is how:
```Go
arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // this creates a 2 dimensional array easily.
fmt.Println(arr[0][0]) // this is how you access elements on a 2 dimensional array.
```

## Maps
Maps are data structure which stores its data in a key-value format. Just like the array, the type of data to be stored for both the key and value must be specified.

Creating a map in golang and some basics of using a map:
```Go
	lawson := map[string]string{"name": "lawson", "animal": "shark", "color": "blue", "location": "ocean"}
	// checking the length/size of the entire map
	fmt.Println(len(lawson))

	// how to access each value of a map
	lawson["color"] = "green" // changes the color of lawson map from blue to green
	lawson["animal"] = "Lion"

	// iterate over a map using the range statement to get the key and value of a map
	for key, value := range lawson {
		fmt.Printf("key[%v]: value[%v]\n", key, value)
	}

	// how to delete a key in map
	delete(lawson, "color")
	fmt.Println(lawson)
	// output: map[animal:Lion location:ocean name:lawson]
	// this removes the key "color from the map"
```
Maps can be used to store multiple values of different types but it is not advisable to do so when creating a map. Here is how to do that in go.
```Go
    anyData := map[any]any{}
	anyData[1] = "lawson"

	fmt.Println(anyData) // output: map[1:lawson]
```
This is not advised as it defeats the concept of golang being statically typed.

## Structs

## Interfaces

## Resources
1. [Understanding arrays and slices - Digital Ocean](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go)
