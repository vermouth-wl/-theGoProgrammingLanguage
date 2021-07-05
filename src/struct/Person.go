package main

import "fmt"

type Person struct {
	Name  string // 姓名
	Birth string //出生日期
	ID    int64  // 身份证号
}

// 指针类型，修改个人信息
func (person *Person) changeName(name string) {
	person.Name = name
}

// 使用非指针类型修改字段
func (person Person) changeName2(name string) {
	person.Name = name
}

// 非指针类型，打印个人信息
func (person Person) printMess() {
	fmt.Printf("我的名字是 %v, 生日是 %v, 我的身份证是 %v\n", person.Name, person.Birth, person.ID)
	// 尝试修改结构体属性，对原结果并无影
}

func main() {
	p := Person{
		"贝尔摩德",
		"2004-02-08",
		500384200402087215,
	}
	p.printMess()       // 打印结构体信息
	p.changeName("苦艾酒") // 使用指针类型修改数据并打印信息
	p.printMess()

	p.changeName2("王磊") // 使用非指针类型修改数据并打印信息
	p.printMess()
}
