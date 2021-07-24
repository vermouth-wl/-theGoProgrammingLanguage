package main

import "fmt"

// 用闭包特性实现一个简单的计数器

func createCounter(initial int) func() int {

	// 传入的initial如果为负数则计0
	if initial < 0 {
		initial = 0
	}

	// 引用initial，创建一个闭包
	return func() int {
		initial++
		// 返回当前计数
		return initial
	}
}

func main() {

	// 计数器1
	c1 := createCounter(1)

	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3

	// 计数器2
	c2 := createCounter(10)

	fmt.Println(c2()) // 11
	fmt.Println(c2()) // 12

	fmt.Println(c1()) // 4

	// 计数器3
	c3 := createCounter(-1)

	fmt.Println(c3()) // 1
}
