package main

import "fmt"

// var strg string = "123"

// func foo(n int) {
// 	if str := "hello world!"; str != "" {
// 		fmt.Println(456)
// 	}
// 	// fmt.Println(str)
// 	fmt.Println(strg)
// }

func Accumulate(n int) func() int {
	value := n
	return func() int {
		value++
		return value
	}
}

func main() {
	// 1. 什么是闭包？
	//  1.1 变量的作用域
	//     局部变量：函数的内部，if，for的块内部
	//     全局变量：包内部
	//     函数的形参：函数的内部
	//  实现原理是什么？
	//     1. go语言是编译语言，不是解释脚本语言
	//     2. 编译、链接、打包 => 函数/二进制码片段/可执行文件
	//     3. 运行阶段：OS启动进程，初始化不同的段 segment，data，bss，code，stack（栈），heap（堆）
	//     4. 函数的调用和返回
	//        函数内部的局部变量，函数的形参，都是在函数调用后，临时分配在stack段上的，一旦退出，内存被销毁（编译器在编译时会发现，并报错）
	//        全局变量：初始化的全局变量分配在data段，未初始化的全局变量分配在bss段中，所以可以在包的任何地方访问。
	//     5. 形参，局部变量，值类型，不用new关键字，通常分配在stack上
	//     6. 用new关键字的引用类型，指针分配在stack上，内存分配在heap
	//  1.2 闭包：函数引用了函数作用域以外的自由变量，通常我们把自由变量所在的环境成为引用环境
	//      因为函数的存在，自由变量不得释放或者删除， 函数+引用环境 = 闭包
	//     闭包的实现原理
	//     7. 凡是闭包中的自由变量，本来应该分配在stack上，但是编译器经过分析后，分配到了heap，所以在函数退出后，还是可以访问这个自由变量
	//        变量的逃逸：go build --gcflags=-m main.go
	//     不同语言处理闭包的情况：
	//     python，JS，C++，Java
	//     python，脚本语言，天然支持闭包
	//     JS，脚本语言，通过作用域链实现
	//     C++，C++11以前，通过stl或者boost模拟，C++11以后增加了原生支持
	//     java JDK1.8以前，使用Function接口和Lambda表达式实现闭包，闭包内的数据是只读的。
	//          JDK1.8以后，可以使用匿名类

	//     8. 每次对立调用会产生不同的闭包。

	// foo(5)
	// fmt.Println(str)

	a := Accumulate(1)
	fmt.Println(a())

	b := Accumulate(10)
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	fmt.Println(a())
}
