package main

import (
	"fmt"
	"time"
	"math/rand"
)

func receiver(c chan int){
	// for {
		// m := <-c
		// fmt.Println("in func()", m)

		// 跟随4. 如果channel存在，可以成功读出数据，对方如果没有发送数据，接收数据阻塞
		//        如果channel关闭了，那么每次读取数据都是失败的，所以此时返回的数据是初始化的0
		// 有两种方法
		//  方法一，
		// m, ok := <- c
		// if !ok {
		// 	break
		// }else{
		// 	fmt.Println("in func()", m)
		// }
	// }

	// 方法二,
	for m := range c {
		fmt.Println("in func()", m)
	}
}

func channelFunc(){
	// 1. channel: 是goroutine用于发送数据和接收数据用的
	//    在任意一个时间点，只能有一个goroutine操作channel
	//    channel在go语言中，是第一公民（函数也是第一公民）
	//    保证数据的先入先出
	//    channel是引用类型，但是需要一个类型修饰
	c := make(chan int)
	go receiver(c)
	// go func(c chan int){
	// 	for {
	// 		m := <-c
	// 		fmt.Println("in func()", m)
	// 	}
	// }(c)

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)

}

// 2. channel可以是有方向的
//    chan<- int(方向)，说明你只可以向这个频道发送消息
//    <-chan int，说明你只可以从这个频道接收消息
func createChannel(id int) chan<- int {
	c := make(chan int)
	go func(){
		for {
			fmt.Printf("Worker No. %d received %c .\n", id, <-c)
		}
	}()
	return c
}

func channelFunc2(){
	var channels [20]chan<- int
	for i := 0; i < len(channels); i++ {
		channels[i] = createChannel(i)
	}

	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

// 3. buffered channel
func bufferedChannel(){
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}
// 4. channel的关闭
func closeChannel(){
	c := make(chan int)
	go receiver(c)
	c <- 1
	c <- 2
	close(c)
	time.Sleep(time.Millisecond)
}

// 5. CSP模型：Communication Sequential Process
// 不要通过共享内存来通信，要通过通信来共享内存

// 6. main函数如何等待所有channel的数据读取完成，再退出？
//   两个channel
func createChannel3(id int) (chan<- int, <-chan bool) {
	c  := make(chan int)
	c2 := make(chan bool)

	go func(c <-chan int, c2 chan<- bool){
		for {
			fmt.Printf("Worker No. %d received %c.\n", id, <- c)
			c2 <- true
		}
	}(c, c2)

	return c, c2
}

func channelFunc3(){
	var channels [20]chan<- int
	var dones    [20]<-chan bool
	for i := 0; i < len(channels); i++ {
		channels[i], dones[i] = createChannel3(i)
	}

	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
		<-dones[i]
	}
}

// 7. 如何提高效率？而不是顺序的？
func createChannel4(id int) (chan<- int, <-chan bool) {
	c  := make(chan int)
	c2 := make(chan bool)

	go func(c <-chan int, c2 chan<- bool){
		for {
			fmt.Printf("Worker No. %d received %c.\n", id, <- c)
			c2 <- true
		}
	}(c, c2)

	return c, c2
}

func channelFunc4(){
	var channels [20]chan<- int
	var dones    [20]<-chan bool
	for i := 0; i < len(channels); i++ {
		channels[i], dones[i] = createChannel4(i)
	}

	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < len(dones); i++ {
		<-dones[i]
	}
}

// 8. 问题升级，如何发两轮收两轮
func createChannel5(id int) (chan<- int, <-chan bool) {
	c  := make(chan int)
	c2 := make(chan bool)

	go func(c <-chan int, c2 chan<- bool){
		for {
			fmt.Printf("Worker No. %d received %c.\n", id, <- c)
			go func(){
				c2 <- true
			}()
		}
	}(c, c2)

	return c, c2
}

func channelFunc5(){
	var channels [20]chan<- int
	var dones    [20]<-chan bool
	for i := 0; i < len(channels); i++ {
		channels[i], dones[i] = createChannel5(i)
	}

	for i := 0; i < len(channels); i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < len(channels); i++ {
		channels[i] <- 'A' + i
	}

	for i := 0; i < len(dones); i++ {
		<-dones[i]
		<-dones[i]
	}
}

// 9. select 通信开关，从不同的channel中获取值
//  如果都阻塞了，等待，直到其中一个可以处理
//  如果有多个可以处理，随机选一个
//  如果没有channel可以处理，并且有default语句，就会执行default

func creakeChannelForSelect() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			c <- i
			i++
		}
	}()

	return c
}
func selectFunc(){
	var c1, c2 chan int = creakeChannelForSelect(), creakeChannelForSelect()
	for {
		select {
		case n := <-c1:
			fmt.Printf("Received from c1 %d .\n", n)
		case n := <-c2:
			fmt.Printf("Received from c2 %d .\n", n)
		}
	}
}



func main() {
	// channelFunc()
	// channelFunc2()
	// bufferedChannel()
	// closeChannel()
	// channelFunc3()
	// channelFunc4()
	// channelFunc5()
	selectFunc()
}