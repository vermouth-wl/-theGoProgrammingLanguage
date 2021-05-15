package main

import "fmt"

// 容器（数组、切片、字典）的遍历
func main() {

	// 数组的遍历
	nums := [...]int{1, 2, 3, 4, 5, 6}
	for k, v := range nums {
		fmt.Printf("k: %d, v: %d\n", k, v)
	}

	fmt.Println()

	// 切片的遍历
	slis := []int{1, 2, 3, 4, 5}
	for k, v := range slis {
		fmt.Printf("k: %d, v: %d\n", k, v)
	}

	fmt.Println()

	// 字典遍历
	tmpMap := map[int]string{
		1: "这是1",
		2: "这是2",
	}
	for k, v := range tmpMap {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}

	fmt.Println()
}
