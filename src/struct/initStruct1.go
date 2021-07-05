package main

import "fmt"

// 初始化内嵌结构体
// 车辆结构的组装和初始化

// Wheel 车轮
type Wheel struct {
	Size int
}

// Engine 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// Car 车
type Car struct {
	Wheel
	Engine
}

func main() {
	// 初始化车辆
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 175,
		},
		// 初始化引擎
		Engine: Engine{
			Type:  "蓝鲸",
			Power: 150,
		},
	}
	fmt.Printf("%+v\n", c)
}
