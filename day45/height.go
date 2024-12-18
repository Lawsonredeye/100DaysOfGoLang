package main

func TreeHeight(t *Tree) int {
	if t == nil {
		return 0
	}
	// left := leftHeight(t.LeftChild)
	// right := rightHeight(t.RightChild)
	// fmt.Println(1 + max(left, right))

	left := TreeHeight(t.LeftChild)
	right := TreeHeight(t.RightChild)
	return 1 + max(left, right)
}

func leftHeight(t *Tree) int {
	if t == nil {
		return 0
	}
	return 1 + leftHeight(t.LeftChild)
}

func rightHeight(t *Tree) int {
	if t == nil {
		return 0
	}
	return 1 + rightHeight(t.RightChild)
}
