package main

import (
	"bufio"
	"fmt"
	"os"
)

func f1() {
	// 1. defer 用于在函数调用结束后调用，用于关闭必须在结束时需要清理的打开资源
	// 2. defer的好处，即使中间出错，defer也会执行，或者代码中间有return，defer也会执行
	// 3. defer是一个栈，后进先出。为什么？
	defer fmt.Println(1)
	defer fmt.Println(1.1)
	defer fmt.Println(1.2)
	defer fmt.Println(1.3)
	fmt.Println(2)
	// panic("error")
	return
	fmt.Println(3)

}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, i)
		// if i == 18 {
		// 	panic("my error")
		// }
	}
}

// 4. defer调用，是在调用的时候确定参数的值
func f2() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	f1()
	writeFile("1.txt")
	f2()
}
