package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 闭包实现生成器
// 创建一个玩家生成器，输入名称，输出生成器
func playerGen(name string) func() (string, int) {
	rand.Seed(time.Now().UnixNano())
	hp := rand.Intn(1000)
	return func() (string, int) {
		// 将变量hp引用到闭包中
		return name, hp
	}
}

func main() {
	// 创建一个玩家生成器
	generator := playerGen("复仇之矛-卡莉斯塔")
	name, hp := generator()
	// 打印玩家及hp
	fmt.Println(name, "-", hp)
}
