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
		t.Value = value
	} else {
		for t.RightChild != nil {
			if value < t.Value {
				t = t.LeftChild
			} else {
				t = t.RightChild
			}
		}
		t.Value = value
	}
}

func (t *Tree) Find(value int) bool {
	// check if the tree is nil
	if t == (&Tree{}) {
		return false
	}

	// if we found it at the root node
	if t.Value == value {
		return true
	}

	for t != nil {
		if t.Value == value {
			return true
		}
		if value < t.Value {
			t = t.LeftChild
		} else {
			t = t.RightChild
		}
	}
	return false
}

func main() {
	// Creating a binary search tree
	// left < parent
	// right > parent
	// [10, 5, 15, 6, 1, 8, 12, 18, 17]

	var tree = &Tree{Value: 5}

	// adding child nodes
	tree.InsertChild(10)
	tree.InsertChild(15)

	// finding values
	fmt.Println(tree.Find(15)) // outputs: true
	fmt.Println(tree.Find(16)) // expect: false
	fmt.Println("bsh")
}
