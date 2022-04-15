package main

import "fmt"

func main() {
	// 1. 与 && 或 || 非 !
	// &&: 双目运算符，只要有一个是false，结果就是false，只有两个都是true，结果才是true
	// b1, b2 := true, false
	// fmt.Println(b1 && b2)
	// b1, b2 = false, true
	// fmt.Println(b1 && b2)
	// b1, b2 = false, false
	// fmt.Println(b1 && b2)
	// b1, b2 = true, true
	// fmt.Println(b1 && b2)

	// || 双目运算符，如果两个值中有一个是true，结果就是true，两个值都是false，结果才是false
	// b1, b2 := true, false
	// fmt.Println(b1 || b2)
	// b1, b2 = false, true
	// fmt.Println(b1 || b2)
	// b1, b2 = true, true
	// fmt.Println(b1 || b2)
	// b1, b2 = false, false
	// fmt.Println(b1 || b2)

	// ! 单目运算符， !true -> false;  !false -> true
	// b1 := true
	// fmt.Println(!b1)
	// b1 = false
	// fmt.Println(!b1)

	// 与或非的优先级不是一样的，非最高，与其次，或最低
	// fmt.Println(1 + 2*3 + 4/5)
	// 1 + 6 + 0 = 7
	// 用数字栈和符号栈处理四则运算
	// 1 + 2 * 3 +
	// 1 + 6 +
	// 7 + 4 / 5
	// 7 + 0
	// 7
	// fmt.Println(true || false && true || !false)
	// true || false && true ||
	// true || false ||
	// true || !false
	// true || true
	// true

	// 2. == !=
	//    > >= < <=
	// s1, s2, s3 := "10", "10", "9"
	// fmt.Println(s1 == s2)
	// fmt.Println(s1 != s3)

	// fmt.Println(s1 > s3)

	// n1, n2, n3 := 10, 10, 9
	// fmt.Println(n1 == n2)
	// fmt.Println(n1 != n3)

	// fmt.Println(n1 > n3)

	// 3. | 双目运算符，按位或
	//    ^ 双目运算法，按位异或: 如果两数字不一样，true，一样，返回false
	// n1, n2 := 5, 11
	// fmt.Println(n1 | n2)

	// fmt.Println(n1 ^ n2)
	//    ^ 单目运算符，取反
	//  如何计算一个负数的值 1，确定符号，2，取反加1
	// fmt.Println(^n1)

	// 4. << 双目运算符 左移
	//    >> 双目运算符 右移
	// n1 := 2
	// fmt.Println(4 << n1)
	// fmt.Println(4 >> n1)

	// 5. & 双目运算符，按位与，A & mask，通过配置mask，置零的位是不关心的位，置一的位是需要知道的位
	//    &^ 双目运算符，按位消除，A &^ mask，通过配置mask，置一的位抹成零，置零的位不动
	// n1 := 5
	// fmt.Println(n1 & 4)

	// fmt.Println(n1 &^ 4)

	// 6. go之简洁
	//    go中没有三目运算符
	//    ++, -- 是语句，不是运算符，++和--是不能出现在表达式里的

	// n := 5
	// // n++
	// b := n++
	// fmt.Println(n, b)

	// 7. + - * / %
	// / %
	n1, n2 := 13, 7
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)

	// 8. go语言中的运算符优先级
	// * / % << >> & &^
	// + - | ^
	// == != < <= > >=
	// &&
	// ||

	fmt.Println(5*2 + 12%5<<2 | 5)
	// 5*2+
	// 10+12%5<<
	// 10+2<<2|
	// 10+8|5
	// 18|5
	// 23

	fmt.Println(5 > 4 || 3 < 2 && 7 > 8 || 4 <= 9)
	// 5>4||
	// true||3<2 &&
	// true||false&&7>8 ||
	// true||false&&false||
	// true||false||
	// true||4<=9
	// true||true
	// true

	// 15 &^ 3
	// 15 ^ 3
	// 5 << 2 + 5

}
