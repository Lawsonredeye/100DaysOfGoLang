package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"

	"github.com/jxskiss/base62"
)

// Create a hash for a string passed as a param and returns a positive number.
func hashFunc(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

// Encode a byte string into a base62 encoding that can be used for
// handling the encoding of a string.
func encode62(b []byte) string {
	result := base62.EncodeToString(b)
	return result
}

func main() {
	url := "https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go"
	h := hashFunc(url)
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, h)
	r := encode62(buf)
	fmt.Println(r)
}
