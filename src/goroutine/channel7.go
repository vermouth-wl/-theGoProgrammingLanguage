package main

import (
	"bufio"
	"fmt"
	"os"
)

func printInput(ch chan string) {
	// 使用for循环从channel中读取数据
	for val := range ch {
		// 读取到EOF结束信号结束循环
		if val == "EOF" {
			break
		}
		fmt.Printf("输入的字符串是 %s\n", val)
	}
}

func main() {
	// 创建一个无缓冲的channel
	ch := make(chan string)
	go printInput(ch)

	// 从命令行读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		ch <- val
		if val == "EOF" {
			fmt.Println("结束输入!")
			break
		}
	}
	defer close(ch)
}
