package main


// creating a QuickSort algorithm

func Insert(arr []int, value int){
	lastElement := len(arr) - 1
	arr[lastElement] = value

	index := len(arr) - 1

	for index > 0 && arr[index] < arr[parent(index)] {
	swap(&heap[index], &arr[parent(index)]
	index = parent(index)
	}
}

func parent(value int) int {
	index := len(arr) / 2
	return 
}

func swap(a, b *int){
	temp := a
	a = b
	b = temp
}
