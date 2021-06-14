package main

import "fmt"

// 结构体实现接口

// Cat 定义接口
type Cat interface {
	// CatchMouse 抓老鼠
	CatchMouse()
}

// Dog 定义接口
type Dog interface {
	// Bark 吠叫
	Bark()
}

// CatDog 定义结构体，使结构体分别实现接口
type CatDog struct {
	Name string
}

// CatchMouse 实现Cat接口方法
func (catDog *CatDog) CatchMouse() {
	fmt.Printf("%v 捉了老鼠并且吃了它\n", catDog.Name)
}

// Bark 实现Dog接口
func (catDog *CatDog) Bark() {
	fmt.Printf("%v 这只狗一直在叫", catDog.Name)
}

func main() {
	catDog := &CatDog{"薇薇"}

	// 声明一个Cat接口，并将catDog指针类型赋值给cat
	var cat Cat
	cat = catDog
	cat.CatchMouse()

	// 声明一个Dog接口，并将catDog指针类型赋值给dog
	var dog Dog
	dog = catDog
	dog.Bark()
}
