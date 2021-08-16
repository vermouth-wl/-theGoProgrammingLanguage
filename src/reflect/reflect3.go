package main

import (
	"fmt"
	"reflect"
)

// 指针与指针指向的元素
func main() {

	// 声明一个空结构体
	type cat struct {
	}

	// 创建cat的实例
	ins := &cat{}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	fmt.Printf("name: '%v' Kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())

	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称与种类
	fmt.Printf("element name: %v, element kind: %v\n", typeOfCat.Name(), typeOfCat.Kind())
}
