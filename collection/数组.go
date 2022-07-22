package main

/**
长度初始化就不会改变
元素类型一样
 初始化长度不同内存分配不一样：>4  --> 堆,否则 栈。
源码：cmd/compile/internal/types/type.go NewSlice()

*/
func main() {
	var empty = []string{}
	var bookings = []string{"zhangsan", "lisi", "wangwu"}

	empty[0] = ""
	bookings[0] = "Nana1"
	bookings[1] = "Nana2"

	i := len(bookings)
	bookings[5] = "index_at_5"

}
