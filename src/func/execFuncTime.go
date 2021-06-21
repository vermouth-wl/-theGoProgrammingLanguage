package main

import (
	"fmt"
	"time"
)

// 计算函数执行时间
func test() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i <= 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时: ", elapsed)
}

func main() {
	test()
}
