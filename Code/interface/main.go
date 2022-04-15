package main

import (
	abc "awesomeProject/Code/interface/pack"
	"fmt"
	"time"
)

type MyPoint struct {
	P *Point
}

func (p *MyPoint) Print() {
	p.P.Print()
}

func (p *MyPoint) Add(x, y int) {
	p.P.X += x
	p.P.Y += y
}

// 定义接口
type Retriever interface {
	Get(url string) string
	// Post(url string) string
}

type RetrieverPost interface {
	Post(url string) string
}

type RetrieverGetAndPost interface {
	Retriever
	RetrieverPost
}

// 使用者
func download(r Retriever) string {
	fmt.Println("in download()")
	return r.Get("http://www.baidu.com")
}

func download2(r RetrieverGetAndPost) string {
	return r.Get("http://www.baidu.com") + r.Post("abc")
}

func main() {
	// 包的封装
	// 1. 变量和函数，都是用CamelCase（驼峰）大驼峰和小驼峰
	// 2. 首字母大写：公有的-》大驼峰
	// 3. 首字母小写：私有的-》小驼峰
	// 4. 每个目录下（直接地）只能有一个包，包名可以和目录名不一致
	// 5. main包下需要有main函数，作为可执行文件的入口

	p := Point{X: 10, Y: 20}
	p.Print()

	// go语言中没有继承，我们如何扩展自己的组件
	// 6. 组合
	p2 := MyPoint{P: &p}
	// p2.P.Print()
	p2.Add(1, 2)
	p2.Print()

	// 7. 自定义类型（别名）, 自定义类型不会自动获得原始类型的方法
	//    go语言不建议大家使用，推荐大家使用：接口+组合来实现扩展组件

	// 8. 接口 go语言对面向对象的拨乱反正，go语言是面向接口的语言
	// 9. Duck Typing：大黄鸭是不是鸭子？
	//    从分类法的角度说，不是
	//    从外部行为的角度上来说，是
	//    分类（继承） vs 标签（接口）
	//    不同语言的支持程度？
	//      Python、JS支持Duck Typing，在运行的时候才知道接口是不是支持
	//      C++的模板，支持Duck Typeing，在编译的时候才能知道是否支持
	//      Java，传入参数必须实现接口，不支持Duck Typing
	//      go语言支持Duck Typing，支持类型检查

	// 10. 接口： 使用者 vs 实现者
	//     使用者：定义接口
	//     实现者：不需要显示声明自己实现了接口，只要真的实现了接口即可
	//     接口被实现的条件：
	//        1. 方法名称一致
	//        2. 参数列表一致
	//        3. 返回参数列表一致
	//        4. 接口的所有方法必须都实现

	var r, r2 Retriever
	r = &abc.MyRetriever{Something: "123"}
	// fmt.Println(download(r))

	r2 = &abc.HTTPRetriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	// fmt.Println(download(r2))

	// 11. 接口如何确定实例的真正类型
	fmt.Printf("%T, %v, %T, %v\n", r, r, r2, r2)

	// 12. 通过类型断言；Type Assertion
	if realR, ok := r.(*abc.MyRetriever); ok {
		fmt.Printf("OK, %s\n", realR.Something)
	} else {
		fmt.Println("Not an *abc.MyRetriever")
	}

	if realR2, ok := r2.(*abc.HTTPRetriever); ok {
		fmt.Printf("OK, %s\n", realR2.UserAgent)
	} else {
		fmt.Println("Not an *abc.HTTPRetriever")
	}

	// 13. 不要使用接口的指针，用指针接收器就能指向对象

	// 14. interface{}，可以用来表示任何类型
	// 练习，实现一个可以放任何对象的队列 Queue Stack
	q := abc.Queue{1}
	q.Push(5)
	q.Push(10)
	fmt.Println(q)
	fmt.Println(q.Shift())
	fmt.Println(q.Shift())
	fmt.Println(q.Shift())
	fmt.Println(q)

	s := abc.Stack{1}
	s.Push(5)
	s.Push(10)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s)

	// 15. 接口的嵌套组合

}
