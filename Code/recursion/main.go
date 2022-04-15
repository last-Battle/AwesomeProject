package main

import "fmt"

func fact(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func step(n int, stepTemp map[int]int) int {
	if n <= 0 {
		return 0
	}

	// 取缓存
	if value, ok := stepTemp[n]; ok {
		return value
	}

	switch n {
	case 1:
		stepTemp[1] = 1
		return 1
	case 2:
		value := 1 + step(1, stepTemp)
		stepTemp[2] = value
		return value
	case 3:
		value := 1 + step(2, stepTemp) + step(1, stepTemp)
		stepTemp[3] = value
		return value
	default:
		value := step(n-3, stepTemp) + step(n-2, stepTemp) + step(n-1, stepTemp)
		stepTemp[n] = value
		return value
	}
}

func doHanoi(n int, src, dest, hpr string) {
	if n > 1 {
		doHanoi(n-1, src, hpr, dest)
		fmt.Printf("Move No. %d dish from %s to %s\n", n, src, dest)
		doHanoi(n-1, hpr, dest, src)
	} else if n == 1 {
		fmt.Printf("Move No. %d dish from %s to %s\n", n, src, dest)
	} else {
		fmt.Println("Error")
	}
}

func main() {
	// 1. 什么是递归？
	// 函数可以调用其他函数，也可以调用自己
	//   如果一个函数，自己调用自己 => 递归
	//   或者，几个函数，互相循环调用 => 递归
	// 动态规划的思想：把复杂问题化简为简单版本的自身问题; 缓存中间结果，用空间换时间

	// 练习1： n的阶乘
	//     5! = 1 * 2 * 3 * 4 * 5
	// f(n): n = 1 => 1
	// 已知f(n-1), f(n) =  f(n-1) * n
	fmt.Println(fact(5))

	// 练习2：台阶走法
	// 有一行台阶，N级，N是正整数，可以一步一级， 一步两级，或者一步三级，一共有多少种走法？
	// 函数名为step
	// N： N = 1 => 1
	//     N = 2 => 一步两级；先走一步，step(1) => 1 + 1 = 2
	//     N = 3 => 一步三级；一步两级，step(1)；一步一级，step(2)； 1 + 1 + 2 = 4
	//     N = 4 => 一步三级, step(1)； 一步两级, step(2)；一步一级，step(3) => 1 + 2 + 4 = 7
	//     ...
	//     N = n => 一步三级, step(n-3)； 一步两级, step(n-2)；一步一级，step(n-1) => step(n-3) + step(n-2) + step(n-1)

	stepTemp := make(map[int]int)
	fmt.Println(step(500, stepTemp), stepTemp)

	// 练习3：汉诺塔
	// 三个柱子，A，B，C柱，N片，初始，所有的片都在一个柱子上，且从小到大罗放，此柱称为源柱，最后将所有的片挪到另外一个柱子上，这个柱子称为目标柱。
	// 要求，1，一次只能移动一个片，且片移动后必须放在柱子上
	//      2，不允许一个片放在比他小的片上
	// doHanoi(片的编号，源柱，目的柱，辅助柱)
	// n: 1 doHanoi(1, A, B, C)
	// n: 2 doHanoi(1, A, C, B), 2 A->B, doHanoi(1, C, B, A)
	// n: 3 doHanoi(2, A, C ,B), 3 A->B, doHanoi(2, C, B, A)
	// ...
	// n: n doHanoi(n-1, A, C, B), n A->B, doHanoi(n-1, C, B, A)

	doHanoi(4, "A", "B", "C")
	// 用缓存改造这个算法，使得时间复杂度降低

}
