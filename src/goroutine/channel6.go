package main

import (
	"fmt"
	"time"
)

// select超时机制处理
func main() {
	ch := make(chan string)
	quit := make(chan bool)

	// 新开一个协程
	go func() {
		for {
			select {
			case name := <-ch:
				fmt.Println("name: ", name)
			case <-time.After(3 * time.Second):
				fmt.Println("无数据接收，超时处理")
				quit <- true
			}
		}
	}()

	users := []string{"神乐", "苦艾酒", "贝尔摩德", "凌之轩家的小贝尔"}
	for _, v := range users {
		ch <- v
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("程序结束")
}
