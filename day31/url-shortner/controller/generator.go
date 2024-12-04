package controller

import (
	"encoding/binary"

	"github.com/jxskiss/base62"
)

// Encodes a unique ID and then returns
func EncodeHash(u uint32) []byte {
	buff := make([]byte, 4)

	binary.LittleEndian.PutUint32(buff, u)
	return buff
}

// Converts a byte slice into a base64 encoded string.
func Shortener(s string) string {
	id := GenerateUniqueID(s)
	encoded := EncodeHash(id)
	return base62.EncodeToString(encoded)
}
