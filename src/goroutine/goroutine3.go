package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go语言竞争状态

var (
	count int32
	wg    sync.WaitGroup
)

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}
