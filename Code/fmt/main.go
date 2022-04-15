package main

import "fmt"

func main() {
	// 1. 布尔值，%t
	// fmt.Printf("%t\n", true)
	// fmt.Printf("%t\n", false)

	// 2. 整数
	// %b，二进制
	// %d，十进制
	// %o，八进制
	// %x，%X 十六进制
	// fmt.Printf("%b\n", 48)
	// fmt.Printf("%d\n", 48)
	// fmt.Printf("%o\n", 48)
	// fmt.Printf("%x\n", 48)

	// %c，字符
	// %d，十进制数字
	// %q，转义字符的字面值
	// fmt.Printf("%c, %d\n", 65, 65)
	// fmt.Printf("a%qb", 10)

	// 3. 浮点数
	// %b float32 8bit阶数，23bit尾数
	//    float64 11bit阶数，52bit尾数

	// var f float32 = 1.5
	// fmt.Printf("%b\n", f)

	// %e, %E 十进制的科学计数法
	// fmt.Printf("%e, %E\n", 235.12, 123.24)

	// %f, %F调整精度和宽度
	// fmt.Printf("%f\n", 235.12)
	// fmt.Printf("%11f\n", 235.12)
	// fmt.Printf("%011f\n", 235.12)
	// fmt.Printf("%011.2f\n", 235.12)

	// %g %G
	// fmt.Printf("%g, %g\n", 123.5, 0.0000000000000000000047)

	// 4. 指针 %p
	// n := 1
	// fmt.Printf("%p", &n)

	// 5. 字符串 %v 实际值 %#v 字面值
	// m := "\n"
	// fmt.Printf("%#v", m)

	// 6. %T 类型

	// 7. 格式化 Printf
	//    普通输出 Print Println （自带一个回车）
	//    Sprintf，返回一个字符串
	//    Fprintf，向io中打印字符串
	// str := fmt.Sprintf("%09.2f", 123.456)
	// fmt.Println(str)

	// 8. 格式化输入
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println("a is: ", a)
}
