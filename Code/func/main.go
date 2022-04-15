package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 计算表达式的值
// func evaluate(a, b int, op string) int {
// 	switch op {
// 	case "+":
// 		return a + b
// 	case "-":
// 		return a - b
// 	case "*":
// 		return a * b
// 	case "/":
// 		return a / b
// 	default:
// 		panic(fmt.Sprintf("unknown operation: %s", op))
// 	}
// }

// 3. 改造
func evaluate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
}

// 4.1 把函数当作参数(传进去)
func apply(a, b float64, op func(float64, float64) float64) float64 {
	// 如何知道传入函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling %s with parameters %f, %f\n", opName, a, b)
	return op(a, b)
}

// 4.2 把函数当作返回值
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// 7. 可变参数列表函数
func sum(ints ...int) int {
	s := 0
	for i := range ints {
		s += ints[i]
	}
	return s
}

func main() {
	// 1. 函数的声明
	// fmt.Println(evaluate(7, 8, "*"))

	// 2. 函数可以返回多个值
	c, d := div(25, 3)
	fmt.Println(c, d)

	// 3. 改造函数evaluate
	if x, err := evaluate(7, 8, "("); err == nil {
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}

	// 4. 把函数作为参数或者返回值
	fmt.Println(math.Pow(2, 10))
	fmt.Println(apply(2, 10, math.Pow))

	// 当前的累加器，与变量现状无关
	adder := adder2(0)
	for i := 0; i < 10; i++ {
		var sum int
		sum, adder = adder(i)
		fmt.Printf("0 +  ... + %d = %d\n", i, sum)
	}
	// 函数式编程： 函数可以作为参数（闭包），函数是一等公民
	// 正统的函数式编程：
	//     第一，不可变性，不能有状态（变量），只有常量和函数
	//     第二，函数只有一个参数
	// 函数式编程很容易并行化

	// 5. go语言不支持重载
	// 6. go语言不支持缺省参数，可选参数
	// 7. 可变参数列表
	fmt.Println(sum(2, 4, 5, 6, 9))
}
