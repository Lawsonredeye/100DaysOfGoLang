package main

import "fmt"

type Tree struct {
	Value      int
	LeftChild  *Tree
	RightChild *Tree
}

// Edge cases:
// tree is empty, leftchild & rightchild are empty
// leftchild & rightchild are not empty so we loop
func (t *Tree) InsertChild(value int) {
	if value < t.Value {
		for t.LeftChild != nil {
			if value < t.Value {
				t = t.LeftChild
			} else {
				t = t.RightChild
			}
		}
		// t
	}
}

func main() {
	// Creating a binary search tree
	// left < parent
	// right > parent
	// [10, 5, 15, 6, 1, 8, 12, 18, 17]

	fmt.Println("bsh")
}
