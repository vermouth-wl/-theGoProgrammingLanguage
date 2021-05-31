package main

import "fmt"

// break用法
func main() {
OutLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OutLoop
			case 3:
				fmt.Println(i, j)
				break OutLoop
			}
		}
	}
}
