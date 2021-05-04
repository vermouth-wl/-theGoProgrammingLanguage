package main

import "fmt"

// 切片的创建与拓容
func main() {

	// 申明两个数组
	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}

	// 对数组进行切片
	sli1 := arr1[0:2]
	sli2 := arr2[2:4]

	fmt.Printf("sli1 指针地址是 %p, 长度为 %d, 容量为 %d, 值是 %v\n", &sli1, len(sli1), cap(sli1), sli1)
	fmt.Printf("sli2 指针地址是 %p, 长度为 %d, 容量为 %d, 值是 %v\n", &sli2, len(sli2), cap(sli2), sli2)

	// 往容量足够的sli1中添加元素
	newSli1 := append(sli1, 5)
	fmt.Printf("newSli1 指针地址是 %p, 长度为 %d, "+
		"容量为 %d, 值是 %v\n", &newSli1, len(newSli1), cap(newSli1), newSli1)
	fmt.Printf("源 arr1 变成 %v\n", arr1)

	// 往容量不足的sli2中添加元素
	newSli2 := append(sli2, 5)
	fmt.Printf("newSli2 指针地址是 %p, 长度为 %d, "+
		"容量为 %d, 值是 %v\n", &newSli2, len(newSli2), cap(newSli2), newSli2)
	fmt.Printf("源 arr2 变成 %v\n", arr2)
}
