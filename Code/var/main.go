package main

import "fmt"

func f() (int, int) {
	return 100, 200
}

func f2() float64 {
	return 3.14
}

type City int

const (
	Beijing City = iota
	Shanghai
	Guangzhou
	Shenzhen
)

func (c City) String() string {
	switch c {
	case Beijing:
		return "Beijing"
	case Shanghai:
		return "Shanghai"
	case Guangzhou:
		return "Guangzhou"
	case Shenzhen:
		return "Shenzhen"
	}

	return "N/A"
}

func main() {
	// 1. 变量的声明
	//    var <name> <type>
	//    变量的命名规则：
	//        小驼峰：所有单词排列，第一个单词首字母小写，以后的单词首字母大写
	//        大驼峰：所有单词首字母大写
	//        匈牙利：属性类型_单词排列，所有单词首字母大写
	//        命名参考网站：https://unbug.github.io/codelf
	//        变量命名：大写字母，小写字母，数字，下划线，并且不能以数字开头？
	//            防止二义性：南京市长江大桥
	//            1E11：到底是变量名还是 科学计数法的 1 x 10^11
	//        变量命名：不可以使用关键字，保留字？
	//            防止二义性

	// var firstVariable int
	// firstVariable = 5
	// fmt.Println(firstVariable)

	// 2. 基本类型
	//    bool
	//    string
	//    int, int8, int16, int32, int64
	//    uint, uint8, uint16, uint32, uint64
	//    byte = uint8
	//    rune = int32
	//    float32, float64
	//    complex64, complex128 复数 3 + 4i
	//    指针

	// var c complex128 = 3 + 4i
	// fmt.Println(cmplx.Abs(c))

	// 欧拉公式
	// var n complex128
	// n = cmplx.Pow(math.E, 1i*math.Pi) + 1
	// fmt.Println(n)

	// C\C++, 如果不初始化，其中数据是随机的 “烫”
	// go，数字： 0 0.0，bool： false，字符串 "", 指针 nil

	// 3. 批量格式
	// var a int = 5
	// var b string = "abc"
	// var c bool = true
	// var (
	// 	a int    = 5
	// 	b string = "abc"
	// 	c bool   = true
	// )
	// fmt.Println(a, b, c)

	// 4. 简短格式（编译器在编译阶段做的）
	//    必须能够从复制的右边确定变量的类型
	//    只能在函数内使用，不可以在包范围内使用。
	// a := 5
	// b := "abc"
	// c := true
	// fmt.Println(a, b, c)

	// d, e := 3, "ccc"
	// fmt.Println(d, e)

	// 简短格式是不能重复使用的
	// f := 5
	// f := 3

	// 简短格式中的变量，要求至少有个一变量是新的，就可以
	// f, e := 10, "ddd"
	// fmt.Println(f, e)

	// 5. 如何交换变量
	// a, b := 5, 7

	// c := a
	// a = b
	// b = c

	// 0101 1011
	// a = a ^ b
	// b = a ^ b
	// a = a ^ b

	// go语言中经常用于交换变量的方法
	// b, a = a, b

	// fmt.Println(a, b)

	// 6. 匿名变量
	// a, _ := f()
	// fmt.Println(a)

	// 7. 常量
	// const pie float64 = 3.1415926
	// fmt.Println(pie)

	// 常量可以没有类型
	// const pie = 3.14
	// fmt.Println(pie)

	// 常量，和变量的区别，是编译器在编译期间直接替换
	// a := pie * 5
	// fmt.Println(a)

	// 常量的赋值，不能是函数，也不能是变量
	// const pie = 2.14 + 1
	// const pie = f2()
	// a := 5
	// const pie = a
	// fmt.Println(pie * 5)

	// C\C++ 宏， Java\C# const，全大写字母表示常量
	// go，首字母大写表明是共有变量，一般不使用全大写

	// 8. 常量的批量形式
	// const (
	// 	pie = 3.1415926
	// 	e   = 2.71818281
	// )

	// 批量形式的省略

	// const (
	// 	a = 1
	// 	b
	// 	c = 5
	// 	d
	// )

	// fmt.Println(a, b, c, d)

	// 9. 枚举类型
	// iota 自增值的枚举类型，编译器在编译期间完成的
	// const (
	// 	sunday    = 0
	// 	monday    = 1
	// 	tuesday   = 2
	// 	wednesday = 3
	// 	thursday  = 4
	// 	firday    = 5
	// 	saturday  = 6
	// )

	// const (
	// 	sunday int = iota
	// 	monday
	// 	tuesday
	// 	wednesday
	// 	thursday
	// 	firday
	// 	saturday
	// )

	// fmt.Println(wednesday)

	// const (
	// 	B = 1 << (10 * iota)
	// 	KB
	// 	MB
	// 	GB
	// 	TB
	// 	PB
	// )

	// fmt.Println(MB, PB)

	// 10. 枚举类型和字符串之间的相互转换

	fmt.Printf("%d, %s", Shanghai, Shanghai)

}
