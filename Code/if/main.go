package main

import (
	"fmt"
)

func main() {
	// if
	// 1. if 没有括号
	n := 5
	if n > 0 {
		fmt.Println("Great!")
	}

	if n < 0 {
		fmt.Println("Path 1")
	} else {
		fmt.Println("Path 2")
	}

	// 2. else必须和大括号在一行
	//    if后的大括号必须和if在一行
	// if n < 0
	// { // 错的
	// 	fmt.Println("Path 1")
	// } else {
	// 	fmt.Println("Path 2")
	// }

	// if n < 0 {
	// 	fmt.Println("Path 1")
	// }
	//  else { // 错的
	// 	fmt.Println("Path 2")
	// }

	// 3. 特殊写法
	//     第一，紧凑
	//     第二，contents，err仅仅在块作用域内有效
	// const filename = "1.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	// const filename = "1.txt"
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	// fmt.Printf("%s\n", contents)

	// switch
	// 4. switch 没有括号, 不需要写break, case后的值不需要是常量
	// a := "hello"
	// switch a {
	// case "hello":
	// 	fmt.Println("1")
	// case "world":
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("3")
	// }

	// 5. 一分支多值
	// a := "Hello"
	// switch a {
	// case "hello", "Hello":
	// 	fmt.Println("1")
	// case "world":
	// 	fmt.Println("2")
	// default:
	// 	fmt.Println("3")
	// }

	// 6. 分支表达式：switch可以没有表达式，case首先匹配为true的，运行
	// 7. fallthrough，用于兼容C\C++

	// 成绩： 0 ~ 100， 60以下：D，60~70：C 70~90：B，90以上：A
	score := 95

	switch {
	case score > 100 || score < 0:
		// fmt.Println("Error")
		panic("Error")
	case score >= 90:
		fmt.Println("A")
		fallthrough
	case score >= 70:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 特点回顾：
	// 1. if，switch没有括号
	// 2. if里可以定义变量
	// 3. switch不需要break
	// 4. 没有三目运算符 ?:
}
