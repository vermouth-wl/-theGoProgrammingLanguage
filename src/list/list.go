package main

import (
	"container/list"
	"fmt"
)

// 列表
func main() {

	// 使用New()函数初始化列表
	tempList := list.New()

	for i := 1; i <= 10; i++ {
		tempList.PushBack(i)
	}

	first := tempList.PushFront(0)
	tempList.Remove(first)

	for l := tempList.Front(); l != nil; l = l.Next() {
		fmt.Print(l.Value, " ")
	}
}
