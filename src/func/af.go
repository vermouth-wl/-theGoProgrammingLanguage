package main

import (
	"flag"
	"fmt"
)

// 使用匿名函数作为map的键值

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	flag.Parse()

	var skill = map[string]func(){
		"姓名": func() {
			fmt.Println("姓名节点")
		},
		"性别": func() {
			fmt.Println("性别节点")
		},
		"年龄": func() {
			fmt.Println("年龄节点")
		},
		"地址": func() {
			fmt.Println("地址节点")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("没有找到skill")
	}
}
