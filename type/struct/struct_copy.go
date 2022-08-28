package main

import "fmt"

// 1 定义
type person struct {
	name  string
	age   int
	email string
}

func main() {
	// 2 初始化
	var p1 = person{"冯大海", 32, "aaa@bb.com"}
	var p2 = p1 // 赋值
	// 3 取值
	p2.name = "flynn"
	fmt.Println(p1) // 不改变
	fmt.Println(p2) // 改变

	var p22 = &p1 // 指针
	p22.name = "flynn"
	fmt.Println(p22)
	fmt.Println(p1) // 也改变
}
