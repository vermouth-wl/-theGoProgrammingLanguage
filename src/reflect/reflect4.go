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
	fmt.Println("我正在跑步，目前速度是: ", hero.Speed)
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

	fmt.Println("------------------------------------------------------------------")

	/**
	* 类型对象reflect.StructField和reflect.Method
	 */
	// 通过NumField 获取结构体字段的数量

	for i := 0; i < typeOfHero.NumField(); i++ {
		fmt.Printf("field' name is %s, type is %s, Kind is %s\n",
			typeOfHero.Field(i).Name,
			typeOfHero.Field(i).Type,
			typeOfHero.Field(i).Type.Kind())
	}
	// 获取名称为Name的成员字段的类型对象
	nameField, _ := typeOfHero.FieldByName("Name")
	fmt.Printf("field' name is %s, type is %s, Kind is %s\n", nameField.Name, nameField.Type, nameField.Type.Kind())

	fmt.Println("------------------------------------------------------------------")

	// 通过MethodField 获取结构体方法数量

	// 声明一个Person接口，并用 Hero 作为接收器
	var person Person = &Hero{}
	// 获取接口Person的类型对象
	typeOfPerson := reflect.TypeOf(person)
	// 打印Person的方法类型和名称
	for i := 0; i < typeOfPerson.NumMethod(); i++ {
		fmt.Printf("方法是 %s, 类型是 %s, 种类是 %s\n",
			typeOfPerson.Method(i).Name,
			typeOfPerson.Method(i).Type,
			typeOfPerson.Method(i).Type.Kind())
	}
	method, _ := typeOfPerson.MethodByName("Run")
	fmt.Printf("方法是 %s, 类型是 %s, 种类是 %s\n", method.Name, method.Type, method.Type.Kind())

	fmt.Println("------------------------------------------------------------------")

	// 判断一个变量的Value是否可寻址, 反射修改字段必须满足指针和Elem()方法两个条件
	name := "小贝尔"
	valueOfName := reflect.ValueOf(name)
	fmt.Printf("name 可以被寻址: %t\n", valueOfName.CanAddr())
	valueOfName = reflect.ValueOf(&name)
	fmt.Printf("name 可以被寻址: %t\n", valueOfName.CanAddr())
	valueOfName = valueOfName.Elem()
	fmt.Printf("name 可以被寻址: %t\n", valueOfName.CanAddr())

	fmt.Println("------------------------------------------------------------------")

	kagura := &Hero{
		Name:  "贝尔摩德",
		Age:   17,
		Speed: 200,
	}

	valueOfKagura := reflect.ValueOf(kagura).Elem()
	valueOfName = valueOfKagura.FieldByName("Name")
	// 判断字段的Value是否可以设定变量值
	if valueOfName.CanSet() {
		valueOfName.Set(reflect.ValueOf("凌之轩家的小贝尔"))
	}
	fmt.Printf("贝尔摩德修改后的名字是 %s\n", kagura.Name)

	fmt.Println("------------------------------------------------------------------")
}
