package main

import (
	"bufio"
	"fmt"
	"os"
)

// dup1 输出标准输入中出现次数大于1的行，前面是次数
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) // 从程序的标准输入进行读取
	for input.Scan() {                  // 每一次调用input.Scan()读取下一行
		if input.Text() == "end" { // 控制循环退出
			break
		}
		counts[input.Text()]++ // 调用input.Text()来获取获取读取到的内容
	}

	// 输出字典
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("重复的值: %s\t 出现次数: %d\n", line, n)
		}
	}
}
