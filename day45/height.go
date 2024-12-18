package main

// Calculates the tree height and return the maximum
// nodes in the tree.
func TreeHeight(t *Tree) int {
	if t == nil {
		return 0
	}
	left := TreeHeight(t.LeftChild)
	right := TreeHeight(t.RightChild)
	return 1 + max(left, right)
}
