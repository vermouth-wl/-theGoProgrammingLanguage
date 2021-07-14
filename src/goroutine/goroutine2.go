package main

import (
	"fmt"
	"time"
)

// 使用匿名函数创建一个协程
func main() {

	go func() {
		var times int
		for {
			times++
			fmt.Println("协程1执行到了: ", times)

			time.Sleep(time.Second)
		}
	}()

	go func() {
		var times int
		for {
			times++
			fmt.Println("协程2执行到了: ", times)

			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scan(&input)
}
