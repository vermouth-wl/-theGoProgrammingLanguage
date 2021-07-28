package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用无缓冲通道模拟4个goroutine之间的接力比赛

// wg3 用来等待程序结束
var wg3 sync.WaitGroup

func main() {
	// 创建一个无缓冲通道
	baton := make(chan int)

	// 为最后一位跑步者将计数+1
	wg3.Add(1)

	// 为第1位跑步者持有接力棒
	go Runner(baton)

	// 开始比赛
	baton <- 1

	// 等待比赛结束
	wg3.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int

	// 等待接力棒
	runner := <-baton

	// 开始围绕跑道跑步
	fmt.Printf("选手 %d 开始跑了\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("选手 %d 到了\n", newRunner)
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗
	if runner == 4 {
		fmt.Printf("选手 %d 跑完了，比赛结束\n", runner)
		wg3.Done()
		return
	}

	// 将接力棒交给下一个跑步者
	fmt.Printf("选手 %d 将接力棒交给了 %d\n", runner, newRunner)

	baton <- newRunner
}
