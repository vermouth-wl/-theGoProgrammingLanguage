package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 使用原子函数来对共享资源进行加锁

var (
	counter int64
	wg1     sync.WaitGroup
)

func main() {
	wg1.Add(3)
	go incCounter(1)
	go incCounter(2)
	go incCounter(3)
	wg1.Wait()
	fmt.Println(counter)
}

func incCounter(id int) {
	defer wg1.Done()
	for count := 0; count < 2; count++ {
		// 安全的对counter + 1
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
