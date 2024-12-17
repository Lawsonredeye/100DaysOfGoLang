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
// func (t *Tree) InsertChild(value int) {
// 	// Check if the tree is an empty tree with no value
// 	if t == nil {
// 		newNode := &Tree{Value: value}
// 		fmt.Println("New node created:", newNode.Value)
// 		t = newNode
// 		return
// 	}

// 	if value < t.Value {
// 		for t.LeftChild != nil {
// 			if value < t.Value {
// 				t = t.LeftChild
// 			} else {
// 				t = t.RightChild
// 			}
// 		}
// 		t.Value = value
// 	} else {
// 		for t.RightChild != nil {
// 			if value < t.Value {
// 				t = t.LeftChild
// 			} else {
// 				t = t.RightChild
// 			}
// 		}
// 		t.Value = value
// 	}
// }

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

	temp := t

	count := 0
	for temp != nil {
		count += 1
		if temp.Value == value {
			return true
		}
		if value > temp.Value {
			fmt.Println("Going Left")
			fmt.Println("leftChild:", temp.LeftChild)
			fmt.Println("value:", temp.Value)
			fmt.Println("rightChild:", temp.RightChild)
			temp = temp.LeftChild
		} else {
			fmt.Println("Going Right")
			temp = temp.RightChild
			fmt.Println("leftChild:", temp.LeftChild)
			fmt.Println("value:", temp.Value)
			fmt.Println("rightChild:", temp.RightChild)
		}
		fmt.Println("Iteration:", count)
	}
	return false
}

// func Find(t **Tree,  value int) bool {
// 	if *t == nil {
// 		return False
// 	}

// 	t.Find
// }

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
	var emptyTree *Tree

	// adding child nodes

	// for _, child := range childValues {
	// 	fmt.Println(child)
	// 	tree.InsertChild(child)
	// }

	InsertChild(&tree, 15)
	InsertChild(&tree, 8)
	InsertChild(&tree, 7)
	InsertChild(&tree, 4)
	InsertChild(&tree, 2)
	InsertChild(&tree, 3)

	// finding values
	fmt.Println("")
	fmt.Println(tree.Find(5)) // expect: true
	fmt.Println()
	fmt.Println(tree.Find(3)) // expect: true

	fmt.Println()
	fmt.Println()
	// Tree traversal
	fmt.Println("Tree Traversal")
	traversePreOrder(tree)

	fmt.Println("Testing for empty tree")
	InsertChild(&emptyTree, 18)

	fmt.Println(emptyTree.Find(18))
}
