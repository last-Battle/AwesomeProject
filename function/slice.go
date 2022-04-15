package function

import "fmt"

func Slice() {
	var numbers = make([]int, 3, 5)
	numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d slice=%v\n", len(x), cap(x), x)
}
