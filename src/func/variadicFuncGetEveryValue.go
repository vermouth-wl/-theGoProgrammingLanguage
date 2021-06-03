package main

import (
	"bytes"
	"fmt"
)

// 通过可变参数（变参函数）获取每一个参数的值

// 定义一个函数，参数数量为0……n,类型限制为string
func joinStrings(slist ...interface{}) string {

	// 定义一个字节缓冲，快速连接字符串
	var b bytes.Buffer

	// 遍历可变参数列表
	for _, s := range slist {
		// 对s进行类型断言，只接收string类型
		switch s.(type) {
		case string:
			// 将遍历出的字符串输入到字节数组中
			b.WriteString(s.(string))
		default:
			fmt.Println("必须传入string类型")
		}
	}
	return b.String()
}

func main() {
	// 输入字符串并连接
	fmt.Println(joinStrings("我", "叫", "贝", "尔", "摩", "德"))
}
