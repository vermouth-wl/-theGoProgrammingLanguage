package main

import (
	"fmt"
	"time"
)

func send(ch chan int, begin int) {
	// 循环向通道发送数据
	for i := begin; i < begin+10; i++ {
		ch <- i
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1, 0)
	go send(ch2, 10)

	// 主goroutine休眠1s，保证调度成功
	time.Sleep(time.Second)

	for {
		select {
		case val := <-ch1:
			fmt.Printf("从协程ch1中获取到 %d\n", val)
		case val := <-ch2:
			fmt.Printf("从协程ch2中获取到 %d\n", val)
		case <-time.After(2 * time.Second): // 超时设置
			fmt.Println("程序超时")
			return
		}
	}
}
