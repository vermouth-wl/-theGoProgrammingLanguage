package main

import (
	"fmt"
	"os"
)

// echo2 输出其命令行参数
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += s + sep + arg
		sep = " "
	}
	fmt.Println(s)
}
