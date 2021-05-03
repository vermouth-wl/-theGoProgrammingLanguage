package main

import (
	"flag"
	"fmt"
)

// 使用flag从命令行中读取参数
func main() {

	// 定义一个类型是string, 名称为surname的命令行参数
	// 参数依次是命令行参数的名称，默认值，提示
	surname := flag.String("surname", "王", "您的姓")

	// 定义一个类型是string, 名称为personalName的命令行参数
	// 除了返回指针类型结果还可以直接传入变量地址获取参数值
	var personalName string
	flag.StringVar(&personalName, "personalName", "小二", "您的名")

	// 定义一个类型是int, 名称为id的命令行参数
	id := flag.Int("id", 0, "您的ID")

	// 解析命令行参数
	flag.Parse()
	fmt.Printf("我是 %v %v, 我的id是 %v\n", *surname, personalName, *id)
}
