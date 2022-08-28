package main

import "fmt"

// 指针是一种数据类型
// 用于表示数据的内存地址

func main() {
	// 声明一个 字符串类型的变量并赋值
	var v1 string = "flynn"
	// 声明一个 字符串的指针类型的变量，并赋值为 name 变量对应的内存地址
	var v2 *string = &v1
	// *指针类型，取值
	fmt.Println(v1, v2, *v2) //

	// 仅仅修改值，地址不会变化
	v1 = "hai"
	fmt.Println(v1, v2, *v2) //
}
