package main

import "fmt"

func changeSlice(s []int) {
	s[0] = 125
}

func main() {
	// 1. slice定义
	//    通过数组来理解slice
	//    在go语言中，数组是连续的内存空间，slice是对数组(或slice)的这样连续空间片段的引用
	//    slice定义：[起始位置:终止位置]半闭半开区间 [起始位置,终止位置)

	// arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	//              0  1  2  3  4  5  6  7
	// slice := arr[2:6]

	// fmt.Println(arr, slice)

	// slice的省略定义
	// slice = arr[2:]
	// fmt.Println("arr[2:]", slice)
	// slice = arr[:6]
	// fmt.Println("arr[:6]", slice)
	// slice = arr[:]
	// fmt.Println("arr[:]", slice)

	// 2. slice的性质：
	//        包括：地址，大小，容量
	//        是引用类型
	// slice = arr[2:6]
	// fmt.Printf("大小: %d, 容量: %d\n", len(slice), cap(slice))

	// changeSlice(slice)
	// fmt.Println(arr, slice)

	// slice2 := arr[4:]
	// changeSlice(slice2)
	// fmt.Println(arr, slice, slice2)
	// 同一个数组可以被不同的slice引用

	// 3. Slice的遍历
	// for _, v := range slice2 {
	// 	fmt.Println(v)
	// }

	// 4. Slice on Slice
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[1:5]
	s2 := s1[3:5]
	fmt.Println(arr, s1, s2)
	// Slice的范围是有cap决定的，而不是len
	// 但是Slice的扩展不可以越界
	// s2 = s1[3:8]
	// fmt.Println(s2)

	// slice的扩展不可以向前
	// s2 = s1[-1:2]
	// fmt.Println(s2)

	// 5. append用于在slice后添加元素
	s3 := arr[5:7]
	s4 := append(s3, 107)
	s5 := append(s4, 108)
	s6 := append(s5, 109, 110, 111)

	fmt.Println(arr)
	fmt.Println("s3:", s3, len(s3), cap(s3))
	fmt.Println("s4:", s4, len(s4), cap(s4))
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("s6:", s6, len(s6), cap(s6))

	// 调用append如果不越界，改写当前数组内值
	// 如果越界，分配新的底层数组，数组大小按照slice的len翻倍分配

	var s7 []int
	// 不初始化，nil
	fmt.Println(s7 == nil)
	for i := 0; i < 100; i++ {
		s7 = append(s7, i)
		// fmt.Printf("len: %d, cap: %d\n", len(s7), cap(s7))
	}

	// slice的基本操作
	// 6. 创建slice
	// 方法一，在已经存在的数组上创建（如上）
	// 方法二，通过直接数组建立
	s11 := []int{1, 2, 3, 4}
	fmt.Printf("type: %T,%T\n", s11, arr)

	// 方法三，只知道长度，内容不知道，用make
	s12 := make([]int, 16)
	fmt.Printf("len: %d, cap: %d\n", len(s12), cap(s12))

	s13 := make([]int, 10, 32)
	fmt.Printf("len: %d, cap: %d\n", len(s13), cap(s13))

	// 7. 拷贝数据
	fmt.Println(s12)
	fmt.Println(copy(s12, s11))
	fmt.Println(s12)

	// 8.删除数据
	// 方法一：append
	s20 := append(s12[:3], s12[4:]...)
	fmt.Println(s20)
	// 方法二：copy
	s20 = s12[:3+copy(s12[3:], s12[4:])]
	fmt.Println(s20)

	// pop, push, shift, unshift, insert, extend
	var s []int
	i := 5

	// 9. push: append
	s = append(s, i)
	s = append(s, i+1)
	s = append(s, i+2)
	fmt.Println(s)

	// 10. pop
	v := s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(v, s)

	// 11. shift
	v = s[0]
	s = s[1:]
	fmt.Println(v, s)

	// 12. unshift: append
	s = append([]int{i + 3}, s...)
	fmt.Println(s)

	// 13. insert: 指定插入后所在的位置
	// s = append(append(s[:1], i+4), s[2:]...) //错的
	s = append(s[:1], append([]int{i + 4}, s[1:]...)...)
	fmt.Println(s)

	// 14. extend
	s = append(s, make([]int, 10)...)
	fmt.Println(s)

}
