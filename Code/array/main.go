package main

import "fmt"

func changeArray(arr [5]int) {
	arr[0] = 125
}

func changeArrayWithPtr(arr *[5]int) {
	arr[0] = 125
}

func changeInt(n int) {
	n = 125
}

func main() {
	// 1. 数组的定义方法：数组的大小必须是在编译时就能确定，常数
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2}
	arr4 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	// 2. 遍历
	// for i := 0; i < len(arr4); i++ {
	// 	fmt.Println(arr4[i])
	// }

	// range 简洁，美观
	for _, v := range arr4 {
		fmt.Println(v)
	}

	// 3. 数组是值类型, 不同长度的数组，即使类型相同，也是不同类型的数组
	changeArray(arr4)
	fmt.Println(arr4)

	a := 10
	changeInt(a)
	fmt.Println(a)

	// 如果想改动，使用指针
	changeArrayWithPtr(&arr4)
	fmt.Println(arr4)

	// 4. 可以用 == 直接比较两个数组的值
	arr5 := [5]int{125, 2, 3, 4, 5}
	fmt.Println(arr4 == arr5)

	// 5. 多维数组: 编译期间能够确定数组的所有维度值
	// 高级语言中，多维数组是低维数组的数组
	// go语言，连续的内存

	var grid [4][5]int
	grid2 := [2][2]int{1: {3, 4}}

	fmt.Println(grid, grid2)

	grid3 := [2]int{1, 2}
	grid4 := [2][2]int{0: {3, 4}, 1: grid3}
	fmt.Println(grid3, grid4)

}
