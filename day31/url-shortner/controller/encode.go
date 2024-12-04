package controller

import "hash/fnv"

// Generates a unique ID from a string passed into it
// and returns a uint23.
func GenerateUniqueID(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}
