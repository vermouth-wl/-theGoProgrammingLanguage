package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 使用互斥锁来创建一个临界区，保证每次只能有一个goroutine进入临界区，避免出现竞态
var (
	counter5 int64
	wg5      sync.WaitGroup
	mutex    sync.Mutex
)

func incCounter5(id int) {
	defer wg5.Done()
	for count := 0; count < 2; count++ {
		// 创建一个临界区，确保每次只能一个goroutine进入临界区
		mutex.Lock()
		{
			value := counter5
			runtime.Gosched()
			value++
			counter5 = value
		}
		// 释放锁，其他goroutine进入临界区
		mutex.Unlock()
	}
}

func main() {

	wg5.Add(2)
	go incCounter5(1)
	go incCounter5(2)

	wg5.Wait()
	fmt.Println(counter5)
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}
