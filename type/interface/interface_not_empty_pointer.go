package main

import "fmt"

/**
type 接口名称 interface{
 方法名称（） 返回值
}

1. 代指类型
2. 约束
*/

type Anima2 struct {
	name string
	age  int
}

type User2 struct {
	name string
	age  int
}

// 非空接口 也可以表示任意类型
type IAnima interface {
	f11() int
}

// 为结构体 User 构造方法
func (u *User2) f11() int {
	return 888
}

// 为结构体 User 构造方法
func (u *Anima2) f11() int {
	return 777
}

// 基于接口的参数，传入实际的结构体也必须有方法。 为什么不像 java 接口 实现呢。
func doingX2(base IAnima) {
	f11 := base.f11()
	fmt.Println(f11)
}

func main() {
	// 传入指针对象
	doingX2(&User2{"flynn", 2})
	doingX2(&Anima2{"flynn", 2})
}
