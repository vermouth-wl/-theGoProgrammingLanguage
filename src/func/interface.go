package main

import "fmt"

// 函数体实现接口

// Printer 定义一个接口
type Printer interface {
	// Print 打印方法
	Print(interface{})
}

// FuncCaller1 将函数转换为类型后实现接口
type FuncCaller1 func(p interface{})

// Print 实现Printer的Print方法
func (funcCaller1 FuncCaller1) Print(p interface{}) {
	// 调用FuncCaller1函数体本身
	funcCaller1(p)
}

func main() {
	var printer Printer
	// 将匿名函数强制转换为FuncCaller1 赋值给printer
	printer = FuncCaller1(func(p interface{}) {
		fmt.Println(p)
	})
	printer.Print("凌之轩家的小贝尔")
}
