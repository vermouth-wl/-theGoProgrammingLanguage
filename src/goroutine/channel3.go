package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用无缓冲通道模拟2个goroutine间的网球比赛

// wg2 用来等待程序结束
var wg2 sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// 创建一个无缓冲通道
	court := make(chan int)

	// 计数器+2，表示要等待两个goroutine
	wg2.Add(2)

	// 启动两个球手
	go player("贝尔", court)
	go player("苦艾酒", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg2.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 函数退出时调用Done,来通知main函数工作已经完成
	defer wg2.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("玩家 %s 赢了\n", name)
			return
		}

		// 选随机数，然后用这个数来判断是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("玩家 %s 丢球了\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数+1
		fmt.Printf("玩家 %s 击球 %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}
