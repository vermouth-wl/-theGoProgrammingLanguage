package main

import (
	"fmt"
	"time"
)

// 使用通道channel实现并发打印

func printer(c chan string) {
	// 开始循环等待接收数据
	for {
		// 从channel中获取一个数据
		data := <-c
		// 将data视为奥特曼时结束
		if data == "" {
			break
		}
		fmt.Println(data)
	}
	// 通知main函数已经结束循环了
	c <- ""
}

func main() {
	// 创建一个channel
	c := make(chan string)
	// 并发执行printer()函数，传入channel
	go printer(c)
	// 创建消息切片
	users := []string{"凌之轩家的小贝尔", "苦艾酒", "贝尔摩德", "神乐", "桔梗", "日暮戈薇", "梦比优斯奥特曼", "奈克瑟斯奥特曼", "奥特曼"}
	// 循环往channel中发送消息
	for _, user := range users {
		c <- user
		time.Sleep(time.Second)
	}
	// 通知printer()结束循环
	c <- ""

	// 等待printer()结束
	<-c
}
