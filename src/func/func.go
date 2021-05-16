package main

import "fmt"

// 函数的使用
// 返回单参数函数使用
func add(a, b int) int {
	return a + b
}

// 返回多参数函数使用
func manyParam(a, b int) (c, d int) {
	c = a + b
	d = a * b
	return
}

func main() {
	fmt.Printf("调用函数 add() 的执行结果为: %v\n", add(1, 3))

	fmt.Println()

	a := 3
	b := 3
	x, y := manyParam(a, b)

	fmt.Printf("调用函数 manyParam() 的执行结果为: %d + %d = %d\n", a, b, x)
	fmt.Printf("调用函数 manyParam() 的执行结果为: %d x %d = %d\n", a, b, y)
}
