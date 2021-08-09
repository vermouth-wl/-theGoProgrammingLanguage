package main

import (
	"fmt"
	"reflect"
)

// 使用反射对常量和结构体进行类型信息获取

// Enum 定义一个Enum类型
type Enum int

// 定义一个常量
const (
	Zero Enum = 0
)

func main() {
	// 声明一个空结构体
	type cat struct {
	}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射对象的类型名称与种类
	fmt.Printf("typeOfCat的类型是 %v, 种类是 %v\n", typeOfCat.Name(), typeOfCat.Kind())
	// 获取常量Zero的反射对象类型及种类
	typeOfA := reflect.TypeOf(Zero)
	fmt.Printf("Zero的类型是 %v, 种类是 %v\n", typeOfA.Name(), typeOfA.Kind())
}
