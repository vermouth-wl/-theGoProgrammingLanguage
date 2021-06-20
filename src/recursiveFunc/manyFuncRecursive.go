package main

import "fmt"

// 多个函数组成递归

// RevSign 函数
func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

// odd 函数
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

// even 函数
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

// main 函数
func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16))
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
}
