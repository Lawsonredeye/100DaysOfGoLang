package main

import (
	"bytes"
	"fmt"
)

func main() {

	data := []byte("example")
	fmt.Println(bytes.Index(data, []byte("amp")))

	fmt.Println()

	datas := []byte("apple, banana, orange")
	parts := bytes.Split(datas, []byte(","))
	fmt.Println(parts)
	// fmt.Println(bytes.Count(data, []byte("e")))

	// byteSliceA := []byte("golang")
	// byteSliceB := []byte("Golang")

	// fmt.Println(bytes.Equal(byteSliceA, byteSliceB))
	// text := []byte("Example")
	// subText := []byte("ample")
	// contains := bytes.Contains(text, subText)
	
	// if contains {
	// 	fmt.Println(contains)
	// 	fmt.Printf("text: %v\nsubText: %v\n", text, subText)
	// } else {
	// 	fmt.Println(contains)
	// }
	// byt := []byte(characterString)

	// fmt.Printf("%s\n",string(byt))
	// fmt.Println(characterString)

}
