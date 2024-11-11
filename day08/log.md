# Bytes and Http Request - Day 08
Today i dove into the world of learning how to use byte and create byte slices and how it helps when parsing data from the response body of a request.
I understood that the response gotten from a request returns a buffer stream which is in a json format that is encoded as a byte which is considerable faster when sending data over a network.

I learned methods which could be used with bytes like:
1. Contains - which returns true if a given byte possess a the character from the bytes slice
```Go
func main() {
 text := []byte("Example")
	subText := []byte("ample")
	contains := bytes.Contains(text, subText)
	
	if contains {
	    fmt.Println(contains)
	 	fmt.Printf("text: %v\nsubText: %v\n", text, subText)
	 } else {
	 	fmt.Println(contains)
	 }
}
```

2. Count() - This returns the Number of occurrence of a byte in a byte slice or a byte.
```Go
func main() {
   data := []byte("example")
   fmt.Println(bytes.Count(data, []byte("e"))) // Output: 2
}
```

3. Equal() - This Method checks if both byte strings have exactly the same properties.
```Go
func main() {
    a := []byte("lawson")
    b := []byte("Lawson") // L is Upper case
    fmt.Println(bytes.Equal(a, b))
    // outputs:
    // false
}
```

There are so much that can be done but this that were touched today was just the surface and i plan on diving deeper into the world of bytes and net/http library.

Resources:
1. [golang bytes](https://www.scalent.io/golang/golang-byte/)
2. [http clients - gobyexample](https://dlintw.github.io/gobyexample/public/http-client.html)

