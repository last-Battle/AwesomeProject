package main

import "fmt"

type Point struct {
	X int
	Y int
}

func createPoint(X, Y int) *Point {
	p := Point{X: X, Y: Y}
	// 此处返回了局部地址变量？

	return &p
}

// func print(p Point) {
// 	fmt.Printf("Point{X=%d, Y=%d}\n", p.X, p.Y)
// }

func (p *Point) print() {
	fmt.Printf("Point{X=%d, Y=%d}\n", p.X, p.Y)
}

func (p *Point) print2() {
	fmt.Printf("print2: %p\n", p)
}

func main() {
	// 结构体
	// 1. 面向对象：封装，继承，多态
	//    封装：封装就是把抽象的数据和对数据进行的操作封装在一起，数据被保存在内部
	//         程序的其他部分只有通过被授权的操作（成员方法），才能对数据进行操作
	//    继承：子类拥有父类的属性和方法，可以重新定义，也可以追加属性和方法。
	//    多态：所有声称支持父类（接口）的环境中，子类实例（接口实例）都可以正常的工作
	// go不支持继承，以及继承多态
	// go只有struct，没有class

	// 2. struct的定义(结构体是值类型)
	// 3. struct的实例化
	// 方法一：声明+赋值
	var p Point
	p.X = 10
	p.Y = 20
	fmt.Println(p)

	// 方法二：new创建指针+赋值
	// 注意：p->X是错的
	//      (*p).X是可以的，不建议，建议写成p.X（go编译器做的语法糖）
	p2 := new(Point)
	p2.X = 30
	p2.Y = 40
	fmt.Println(p2)

	// 方法三：直接量
	p3 := Point{X: 50, Y: 55}
	fmt.Println(p3)

	//  5. go语言中没有构造函数，可以实现一个工厂函数
	p4 := createPoint(80, 90)
	fmt.Println(p4)

	//  6. 变量到底是在栈上还是在堆上？
	//  C\C++，声明在栈上，new在堆上，在堆上申请的内存需要自己释放
	//  Java，.Net等：只要是引用类型都声明在堆上，所有在对上的对象由gc来释放
	//  go：声明的变量不一定在栈上，有些情况下会逃逸到堆上
	//      凡是分配在堆上以及逃逸到堆上的变量，都由gc来释放
	//      变量的逃逸，由编译器来做决定，程序员职能做逃逸分析
	//          1. 函数的返回值，是通过变量逃逸来实现的
	//          2. 函数返回的引用，是通过变量逃逸实现的。
	//          3. 闭包，通过变量逃逸实现

	//  7. 函数的接收器
	//     接收器可以是指针类型，也可以是非指针类型
	//     go语言不需要this或者self指针，名字自己随便写
	//     接收器是指针类型时，nil也可以调用方法
	// print(p)
	p.print()

	var p5 *Point = nil
	p5.print2()

	//  8. 接收器？用指针类型？还是非指针类型？
	//     原则：如果需要修改内容，就必须使用指针接收器
	//           如果结构体的内存过大，应该尽量使用指针接收器
	//           一致性，应该尽量避免把指针接收器和非指针接收器混合使用。
	// 非指针接收器，是go语言特有的。对调用者透明

}
