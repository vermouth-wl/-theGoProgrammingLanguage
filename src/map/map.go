package main

import "fmt"

// 字典
func main() {

	classMeta1 := make(map[string]string)

	// 添加映射关系
	classMeta1["name"] = "王磊"
	classMeta1["address"] = "重庆市南川区"
	classMeta1["年龄"] = "137亿岁（现宇宙同龄）"

	// 根据key获取Value
	fmt.Printf("name is %v\n", classMeta1["name"])

	// 判断某个键是否存在在于map中
	name, ok := classMeta1["name"]
	if ok {
		fmt.Printf("%s 存在于字典中。\n", name)
	}
}
