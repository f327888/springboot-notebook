package main

import "fmt"

// 结构体做参数和返回值时，在执行时都会被copy 一份。
// 如果不想被 copy, 返回结构体的指针类型。

type MyInt int

// 只有这个类型才能执行
// 可以是类型 可以是指针：指针同一个地址
// 不想用对象 可以用下划线 _
func (i *MyInt) dosome(a1 int, a2 int) int {
	return a1 + a2 + int(*i)
}
func main() {
	var v1 MyInt = 1
	dosome := v1.dosome(11, 22)
	fmt.Println(dosome)

}
