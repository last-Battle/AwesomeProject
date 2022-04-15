package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 1. goroutine 并发编程 并发vs并行
	// go + 函数调用，开协程
	// 什么是协程？什么是线程？
	// 线程是操作系统层次支持的并发，内存共享
	// goroutine，Go语言中的轻量级的线程实现，由Go运行时（runtime）管理。
	//    go程序会智能的将goroutine中的任务分配给每个CPU
	// 2. 特点：
	//    第一，轻量级
	//    第二，非抢占式多任务，协程需要主动交出控制权
	//    第三，编译器级别的并发
	//    第四，多个协程会被分配到一个或者多个线程上运行
	sum := 0

	for i := 0; i < 1000000; i++ {
		go func(i int) {
			for {
				// fmt.Printf("goroutine print something: %d\n", i)
				sum += i
				runtime.Gosched()
			}
		}(i)
	}

	// 3. 竞争冒险， race condition
	// go run -race main.go

	// 4. 子程序调用其实是协程的一个特例（《The Art of Computer Programming Vo.1》）

	// 5. 关于协程不同语言的支持情况：
	// C++： Boost支持
	// Java：不支持
	// Python：支持 yeild，3.5后：async def
	// JavaScript：单线程，yeild，Promise支持异步

	// 6. goroutine和线程的关系，不用程序员来控制，go有自己的调度器
	//    加 go就能把函数送给调度器运行。
	//    在go语言中，函数是否是异步的是不加区别的，Python需要
	//    调度器在什么时间点上切换：
	//        I/O, select, channel, 锁，函数的调用，runtime.Gosched()

	// 7. 查看运行时的线程数

	time.Sleep(time.Millisecond * 10000)
	fmt.Println("end")

}
