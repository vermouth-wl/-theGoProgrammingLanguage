package main

import (
	"fmt"
	"time"
)

// 使用普通函数构建一个协程
func running1() {

	var times int

	// 构建一个无限循环，times每隔1秒+1
	for {
		times++
		fmt.Println("协程1执行到了: ", times)

		// 延时1秒
		time.Sleep(time.Second)
	}
}

// 使用普通函数构建一个协程
func running2() {

	var times int

	// 构建一个无限循环，times每隔1秒+1
	for {
		times++
		fmt.Println("协程2执行到了: ", times)

		// 延时1秒
		time.Sleep(time.Second)
	}
}

// main
func main() {

	// 并发执行程序
	go running1()
	go running2()

	// 接受命令行输入，不做任何处理
	var input string
	fmt.Scan(&input)

}
