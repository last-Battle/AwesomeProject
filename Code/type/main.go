package main

import "fmt"

func main() {
	// 1. bool
	// var a bool
	// a = true

	// var b bool = true

	// var c = true

	// d := true

	// var e bool

	// fmt.Println(a, b, c, d, e)

	// 2. float32 float64
	// a := 3.14 // float64
	// fmt.Printf("a is type of %T\n", a)

	// var b float64
	// fmt.Println(b)

	// 3. byte ASCII
	// var (
	// 	a byte = 97
	// 	b      = 65
	// 	c      = 48
	// )

	// fmt.Printf("%c, %d, %c, %d, %c, %d\n", a, a, b, b, c, c)

	// a = 'a'
	// b = 'A'
	// c = '0'

	// fmt.Printf("%c, %d, %c, %d, %c, %d\n", a, a, b, b, c, c)

	// fmt.Printf("%c\n", 'a'-32)
	// fmt.Printf("%c\n", 'A'+32)

	// fmt.Printf("%d\n", '9'-'0')

	// 4. string
	// var a string = "abcd"
	// fmt.Printf("%T, %s\n", a, a)

	// // 求长度 len() 占内存的长度，不是字符的个数
	// fmt.Printf("length: %d\n", len(a))

	// b := "隔壁王校长"
	// fmt.Printf("length: %d\n", len(b))

	// 字符和字符串的关系
	// go内部的编码格式 — utf-8
	// javaScript，Java Unicode

	// 解释型字符串和非解释型字符串
	// 非解释型字符串，不转译字符，不受换行影响
	str1 := "abcd\n"
	//       01234
	str2 := `abcd\n`
	fmt.Println(str1, str2)

	divDOM := `
	<div>
		<div>
			<img />
		</div>
		<p>
		</p>
	</div>`
	fmt.Println(divDOM)

	fmt.Printf("%c", str1[2])
}
