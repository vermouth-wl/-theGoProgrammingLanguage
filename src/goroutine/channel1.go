package main

import (
	"fmt"
	"time"
)

// 通道channel的使用
func main() {

	// 构建一个通道
	ch := make(chan string)

	users := []string{"苦艾酒", "凌之轩家的小贝尔", "贝尔摩德"}

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("开始协程")
		// 往main的goroutine中发送user信息
		for _, user := range users {
			ch <- user
			time.Sleep(time.Second)
		}
		fmt.Println("结束协程")
	}()

	fmt.Println("等待协程")
	// 接收通道中的值
	for data := range ch {
		fmt.Println(data)
		if data == "凌之轩家的小贝尔" {
			fmt.Println("成功接收到了凌之轩家的小贝尔")
		}
		if data == "贝尔摩德" {
			break
		}
	}
	fmt.Println("接收完成")
}
