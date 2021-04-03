package main

import "fmt"

// 求两个数的最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	fmt.Println(x, y)
	return x
}

// 计算斐波那契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {

	fmt.Println("调用gcd() 函数求得最大公约数是: ", gcd(100, 23))
	fmt.Println("调用fib() 函数求得菲波那切数列的第n个数是: ", fib(9))

}
