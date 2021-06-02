package main

import "fmt"

// Accumulate 闭包的记忆效应
// 提供一个值，每次调用值会对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		value++
		return value
	}
}

func main() {
	// 创建一个累加器，初始值为1
	acc := Accumulate(1)
	// 累加1并打印
	fmt.Println(acc())
	fmt.Println(acc())
	fmt.Println(acc())
	// 打印累加器的函数地址
	fmt.Printf("%p", &acc)

	// 创建一个累加器，初始值为10
	acc2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(acc2())
	fmt.Println(acc2())
	fmt.Println(acc2())
	// 打印累加器函数地址
	fmt.Printf("%p", &acc2)
}
