package main

import "fmt"

// defer延迟语句
func main() {
	fmt.Println("defer开始")
	// 将defer放入延迟栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("defer结束")
}
