package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "This is a bare string"

	hashedString := sha256.New()

	hashedString.Write([]byte(str))

	bs := hashedString.Sum(nil)

	fmt.Printf("%x\n", bs)
}
