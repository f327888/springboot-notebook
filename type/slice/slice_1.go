package main

import "fmt"

func main() {

	// 一般用make ,尽量一次定下容量，减少扩容
	var usrs = make([]int, 1, 3)
	fmt.Println(usrs) // [0]

	// 容易混淆的 [] 容量是0
	// 极端情况
	var v1 = new([]int) // 指针类型，指向的是切面 (有个初始化的概念)
	var v2 *[]int       // 指针类型，指向的是 nil , 而不是切片
	fmt.Println(v1)     // &[]
	fmt.Println(v2)     // <nil>

	var v13 = new([3]int) // 指针类型，指向的是切面 (有个初始化的概念)
	var v23 *[3]int       // 指针类型，指向的是 nil , 而不是切片
	fmt.Println(v13)      // &[0 0 0]
	fmt.Println(v23)      // <nil>

}
