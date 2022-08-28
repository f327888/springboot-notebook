package main

import "fmt"

// 指针是一种数据类型
// 用于表示数据的内存地址

func main() {
	// 声明一个 字符串类型的变量并赋值
	var name string = "flynn"
	//
	var p1 *string = &name
	// 注意 string 类型前面的 *
	var p2 **string = &p1
	var p3 ***string = &p2

	fmt.Println(name) // 空
	fmt.Println(p1)   // <nil>
	fmt.Println(p2)   // 0
	fmt.Println(p3)   // <nil>

}
