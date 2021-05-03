package main

import "fmt"

// go语言指针
func main() {
	str := "我的名字是贝尔摩德"
	strPtr := &str

	fmt.Printf("str的类型是 %T, 值是 %v, 内存地址是 %p\n", str, str, &str)
	fmt.Printf("strPtr的类型是 %T, 值是 %v\n", strPtr, strPtr)

	newStr := *strPtr // 获取指针变量对应的值
	fmt.Printf("newStr的类型是 %T, 值是 %v, 内存地址是 %p\n", newStr, newStr, &newStr)

	*strPtr = "你的名字是苦艾酒" // 通过指针对变量进行赋值
	fmt.Printf("*strPtr的类型是 %T, 值是 %v, 内存地址是 %p\n", *strPtr, *strPtr, strPtr)
	fmt.Printf("str的类型是 %T, 值是 %v, 内存地址是 %p\n", str, str, &str)

}
