package main

import "fmt"

// Factorial 数字阶乘
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i int = 10
	fmt.Printf("数字 %d 经过阶乘后的结果是: %d\n", i, Factorial(uint64(i)))
}
