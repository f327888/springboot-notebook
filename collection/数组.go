package main

import "fmt"

/**
长度初始化就不会改变
元素类型一样
 初始化长度不同内存分配不一样：>4  --> 堆,否则 栈。
源码：cmd/compile/internal/types/type.go NewSlice()
// 声明
 var var_name [size]var_type
// 声明+初始化
 var var_name = [size]var_type{}
*/
func main() {
	var empty = []string{}
	var bookings = []string{"zhangsan", "lisi", "wangwu"}

	empty[0] = ""
	bookings[0] = "Nana1"
	bookings[1] = "Nana2"

	i := len(bookings)
	bookings[5] = "index_at_5"

	// 类型是同一个，并且和声明的类型一致
	var bookings = [50]string{"aa", "bb", "cc"}
	//var bookings = [50]string{1} // 元素的类型不一致错误
	//_______________
	//|_|_|_|_|_|_|________|

	var bookings2 [50]string
	bookings2[0] = "aa1"
	bookings2[1] = "aa2"
	bookings2[2] = "bb2" + "bb3"

	fmt.Printf("The whole array: %v ", bookings2)
}
