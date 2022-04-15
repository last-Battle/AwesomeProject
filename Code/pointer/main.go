package main

import "fmt"

func swapByValue(a, b int) {
	fmt.Printf("in func, will swap by value: %d, %p, %d, %p\n", a, &a, b, &b)
	temp := a
	a = b
	b = temp
	fmt.Printf("in func, swapped by value: %d, %p, %d, %p\n", a, &a, b, &b)
}

func swapByPointer(a, b *int) {
	fmt.Printf("in func, will swap by P: %d, %p, %d, %p\n", *a, a, *b, b)
	temp := *a
	*a = *b
	*b = temp
	fmt.Printf("in func, will swap by P: %d, %p, %d, %p\n", *a, a, *b, b)
}

func main() {
	// 1. 地址(指针 Pointer)
	x, y := 3, 5
	fmt.Printf("%d, %p, %d, %p\n", x, &x, y, &y)

	// 2. 指针变量的创建
	var ptr1 *int = &x
	ptr2 := &y
	fmt.Printf("%p, %T, %p, %d\n", ptr1, ptr1, ptr2, *ptr1)

	// 3. new()
	ptr3 := new(int)
	*ptr3 = 10
	fmt.Printf("%p, %T, %d\n", ptr3, ptr3, *ptr3)

	// 4. 函数参数的三种传递方式
	//    值传递
	//    值类型的引用传递
	//    引用传递

	// 变量交换
	swapByValue(x, y)
	fmt.Println(x, y)

	swapByPointer(&x, &y)
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

	// 与C\C++的不同点：
	// 不同类型的指针不能相互转化，例如， *int32, *float64
	// 任何普通指针类型*T和uintptr之间不能相互转化
	// 指针变量不能参与运算，比如++，--

}
