package main

import (
	"fmt"
	"time"
)

func consume(ch chan int) {
	// 线程最后休息100s后再从channel中读取数据
	time.Sleep(time.Second * 100)
	<-ch
}
func main() {
	// 创建一个长度为2的channel
	ch := make(chan int, 2)
	go consume(ch)

	ch <- 0
	ch <- 1
	//发送数据不被阻塞
	fmt.Println("我空闲了")
	ch <- 2
	fmt.Println("我不能等待100s")

	time.Sleep(time.Second)
}
