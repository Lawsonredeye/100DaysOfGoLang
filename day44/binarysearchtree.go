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
	// Check if the tree is an empty tree with no value
	if t == nil {
		newNode := Tree{Value: value}
		t = &newNode
		return
	}

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
	if t == nil {
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

func traversePreOrder(t *Tree) {
	// Root -> Left -> Right
	if t == nil {
		// fmt.Println("Root:", t.Value)
		return
	}

	fmt.Println("value:", t.Value)
	traversePreOrder(t.LeftChild)
	traversePreOrder(t.RightChild)
}

func main() {
	// Creating a binary search tree
	// left < parent
	// right > parent
	// childValues := []int{10, 15, 6, 1, 8, 12, 18, 17}

	var tree = &Tree{Value: 5}
	var emptyTree Tree

	// adding child nodes

	// for _, child := range childValues {
	// 	fmt.Println(child)
	// 	tree.InsertChild(child)
	// }

	tree.InsertChild(15)
	tree.InsertChild(8)
	tree.InsertChild(7)
	tree.InsertChild(4)
	tree.InsertChild(2)
	tree.InsertChild(3)

	// finding values
	fmt.Println(tree.Find(2)) // expect: true
	fmt.Println(tree.Find(3)) // expect: false

	// Tree traversal
	traversePreOrder(tree)

	fmt.Println("Testing for empty tree")
	emptyTree.InsertChild(18)

	fmt.Println(emptyTree.Find(17))
}
