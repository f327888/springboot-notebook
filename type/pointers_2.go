package main

import "fmt"

// 指针是一种数据类型
// 用于表示数据的内存地址

func main() {
	// 声明一个 字符串类型的变量（默认初始值是空）
	var v1 string
	// 声明一个 字符串的指针类型的变量（默认初始化值是 nil）
	var v2 *string

	var v3 int
	var v4 *int

	// 声明一个 字符串类型的变量并赋值
	var name string = "flynn"
	// 声明一个 字符串的指针类型的变量，并赋值为 name 变量对应的内存地址
	var pointer *string = &name

	fmt.Println(v1)      // 空
	fmt.Println(v2)      // <nil>
	fmt.Println(v3)      // 0
	fmt.Println(v4)      // <nil>
	fmt.Println(pointer) // 0xc000088220

}
