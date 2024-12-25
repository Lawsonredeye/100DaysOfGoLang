package main
import "fmt"


// Min-Heap Implementation
func Insert(arr []int, value int) []int {
	arr = append(arr, value)

	index := len(arr) - 1

	for index > 0 && arr[index] < arr[parent(index)] {
	swap(&arr[index], &arr[parent(index)])
	index = parent(index)
	}

	return arr
}

func parent(value int) int {
	index := (value - 1) / 2
	return index
}

func swap(a, b *int){
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	arr := make([]int, 0)

	array := []int{9,2,5,1,6,8}

	for _, value := range array {
		arr = Insert(arr, value)
	}

	fmt.Println(arr)
}
