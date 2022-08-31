package main

import "fmt"

// 结构体做参数和返回值时，在执行时都会被copy 一份。
// 如果不想被 copy, 返回结构体的指针类型。

type Base struct {
	name string
}

type Son struct {
	Base // 匿名的方式，如果改为 base Base,则无法继承Base 的方法
	age  int
}

// Base 结构体的指针类型的的方法
func (b *Base) m1() int {
	return 666
}

// Son 结构体的指针类型的方法
func (s *Son) m2() int {
	return 999
}
func (s *Son) m3() int {
	return 777
}

func main() {
	son := Son{age: 18, Base: Base{
		name: "冯大海",
	}}
	// 可以调用 Base 里面的方法
	result1 := son.m1()
	result3 := son.m3()
	result2 := son.m2()

	fmt.Println(result1, result2, result3)

}
