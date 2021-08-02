package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fix(x int) int {
	if x < 2 {
		return x
	}
	return fix(x-1) + fix(x-2)
}

func main() {
	go spinner(100 * time.Millisecond)
	const n = 100
	fibN := fix(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
