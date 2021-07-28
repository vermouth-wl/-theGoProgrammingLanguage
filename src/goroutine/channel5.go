package main

import "fmt"

// 无缓冲通道与有缓冲通道的区别
func main() {

	// 创建一个3个元素缓冲大小的string类型通道
	ch1 := make(chan string, 3)
	// 查看当前通道大小
	fmt.Println(len(ch1))
	// 发送3个string型元素到通道
	ch1 <- "凌之轩家的小贝尔"
	ch1 <- "贝尔摩德"
	ch1 <- "苦艾酒"
	// 查看当前通道的大小
	fmt.Println(len(ch1))

	// 创建一个无缓冲的string类型的通道
	ch2 := make(chan string)
	// 发送string型元素到通道
	ch2 <- "王磊"
	ch2 <- "神乐"
	// 接收数据，无缓冲不接收会报错
	for data := range ch2 {
		fmt.Println(data)
		if data == "" {
			break
		}
	}
}
