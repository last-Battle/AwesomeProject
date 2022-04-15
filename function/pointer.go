package function

import "fmt"

func Pointer() {
	a := 20
	var ip *int

	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	const MAX int = 3

	//b := []int{10, 100, 1000}
	//var i int
	//for i = 0; i < MAX; i++ {
	//	fmt.Printf("a[%d]=%d\n", i, b[i])
	//}
	b := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &b[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("b[%d] = %d\n", i, *ptr[i])
	}

}
