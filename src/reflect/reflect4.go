package main

import (
	"fmt"
	"reflect"
)

// Person 定义一个人的接口
type Person interface {

	// SayHello 和人说你好
	SayHello(name string)

	// Run 跑步
	Run() string
}

type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("你好: ", name, ", 我是: ", hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("我正在跑步，目前速度是: ", string(hero.Speed))
	return "Running"
}

func main() {
	a := Hero{
		Name:  "贝尔",
		Age:   17,
		Speed: 50,
	}
	a.SayHello("苦艾酒")
	a.Run()

	typeOfHero := reflect.TypeOf(a)
	fmt.Printf("Hero的类型是 %s, 种类是 %s\n", typeOfHero, typeOfHero.Kind())
	fmt.Printf("*Hero的类型是 %s, 种类是 %s\n", reflect.TypeOf(&Hero{}), reflect.TypeOf(&Hero{}).Kind())

	typeOfPtrHero := reflect.TypeOf(&Hero{})
	fmt.Printf("*Hero的类型是 %s, 种类是 %s\n", typeOfPtrHero, typeOfPtrHero.Kind())

	// 通过Elem()方法获取指针typeOfPtrHero指向变量的真实类型
	typeOfHero = typeOfPtrHero.Elem()
	fmt.Printf("typeOfPtrHero elem到 typeOfHero, Hero的类型是 %s, 种类是 %s\n", typeOfHero, typeOfHero.Kind())

}
