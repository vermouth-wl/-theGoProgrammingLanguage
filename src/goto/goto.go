package main

import "fmt"

// goto语句
func main() {

	// 不使用goto语句的情况

	var breakAgain bool
	//外循环
	for x := 0; x < 10; x++ {
		// 内循环
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 设置退出标记
				breakAgain = true
				// 退出本次循环
				break
			}
		}
		// 根据标记，还需要借宿一次循环
		if breakAgain {
			break
		}
	}
	fmt.Println("done")
}
