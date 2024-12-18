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
		return
	}

	if value < t.Value {
		if t.LeftChild == nil {
			t.LeftChild = &Tree{Value: value}
			return
		} else {
			t.LeftChild.InsertChild(value)
		}
	} else if value > t.Value {
		if t.RightChild == nil {
			t.RightChild = &Tree{Value: value}
			return
		} else {
			t.RightChild.InsertChild(value)
		}
	}

}

func InsertChild(t **Tree, value int) {
	if *t == nil {
		// create a new node if tree is nil
		newNode := Tree{Value: value}
		fmt.Println("Added new node:", newNode.Value)
		*t = &newNode
		return
	}

	if value > (*t).Value {
		InsertChild(&(*t).LeftChild, value)
	} else {
		InsertChild(&(*t).RightChild, value)
	}
}

func (t *Tree) Find(value int) bool {
	// check if the tree is nil
	if t == nil {
		return false
	}

	// creating a temporary variable for traversal
	temp := t

	count := 0
	for temp != nil {
		count += 1
		if temp.Value == value {
			return true
		}
		if value > temp.Value {
			temp = temp.LeftChild
		} else {
			temp = temp.RightChild
		}
		fmt.Println("Iteration:", count)
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

// Traverse a tree using the LEFT -> ROOT -> RIGHT
func traverseInOrder(t *Tree) {
	if t == nil {
		return
	}

	traverseInOrder(t.LeftChild)
	fmt.Println("Value:", t.Value)
	traverseInOrder(t.RightChild)
}

func traversePostOrder(t *Tree) {
	if t == nil {
		return
	}

	traversePostOrder(t.LeftChild)
	traversePostOrder(t.RightChild)
	fmt.Println("Value:", t.Value)
}

func main() {
	// Creating a binary search tree
	// left < parent
	// right > parent
	childValues := []int{4, 9, 1, 6, 8, 10}

	var tree = &Tree{Value: 7}
	// var emptyTree *Tree

	// adding child nodes

	for _, child := range childValues {
		InsertChild(&tree, child)
		// tree.InsertChild(child)
	}

	// finding values
	fmt.Println("")
	fmt.Println("Finding the value of 5")
	fmt.Println(tree.Find(5)) // expect: true
	fmt.Println()

	fmt.Println("Finding the value of 6")
	fmt.Println(tree.Find(6)) // expect: true

	fmt.Println()
	fmt.Println()
	// Tree traversal
	fmt.Println("Pre Traversal")
	traversePreOrder(tree)

	fmt.Println("InOrder Traversal")
	traverseInOrder(tree)

	fmt.Println("PostOrder Traversal")
	traversePostOrder(tree)

	fmt.Println("Tree Height")
	value := TreeHeight(tree)
	fmt.Println(value)
	// fmt.Println("Testing for empty tree")
	// InsertChild(&emptyTree, 18)

	// fmt.Println(emptyTree.Find(18))
}
