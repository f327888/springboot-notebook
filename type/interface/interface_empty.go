package main

import "fmt"

/**
type 接口名称 interface{
 方法名称（） 返回值
}

1. 代指类型
2. 约束
*/

type Person struct {
	name string
	age  int
}

// 空接口 任意类型
type Base interface {
}

// 参数任意类型
func doX(arg interface{}) {
	fmt.Println(arg)
	// .(type) 转换类型 --> 泛型
	person, ok := arg.(Person)
	if ok {
		fmt.Println(person.name, person.age)
	} else {
		fmt.Println("转换失败")
	}

	switch tp := arg.(type) {
	case Person:
		fmt.Println(tp.name, tp.age)
	case int:
		fmt.Println(tp)
	case string:
		fmt.Println(tp)
	default:
		fmt.Println(tp)
	}

}

func main() {

	dataList := make([]Base, 0) // 可以简写为下面
	//i := make([]interface{}, 0)

	dataList = append(dataList, "大海")
	// 多种类型
	doX("flynn")
	doX(666)
	doX(nil)
	doX(Person{"flynn", 2})
}
