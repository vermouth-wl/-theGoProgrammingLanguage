package main

import "fmt"

// 可变参数（变参函数）

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		// 判断类型（类型断言）
		switch arg.(type) {
		case int:
			fmt.Println(arg, "是一个int类型的值")
		case string:
			fmt.Println(arg, "是一个string类型的值")
		case bool:
			fmt.Println(arg, "是一个bool类型的值")
		case float32:
			fmt.Println(arg, "是一个float32类型的值")
		case float64:
			fmt.Println(arg, "是一个float64类型的值")
		case complex128:
			fmt.Println(arg, "是一个complex64类型的值")
		default:
			fmt.Println("未知类型")
		}
	}
}

func main() {
	var v1 = 123
	var v2 = "王磊"
	var v3 = true
	var v4 = 55.5
	var v5 = 6666666.6
	var v6 = 20 + 10i
	myPrintf(v1, v2, v3, v4, v5, v6)
}
