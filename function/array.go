package function

import "fmt"

func Array() {
	//var a [10]float32
	//b := [10]float32{12412.00, 123123.23}
	//c := [...]float32{123, 123123}
	//d := [5]float32{1: 2.0, 2: 3.0}
	//d[4] = 100
	//fmt.Println(a, b, c, d)
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("前两个元素为：")
	fmt.Println(values[0][0:2])
}
func GetAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
