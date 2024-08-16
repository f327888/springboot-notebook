package main

import "fmt"

/**
type 接口名称 interface{
 方法名称（） 返回值
}

1. 代指类型
2. 约束
*/

type Anima struct {
	name string
	age  int
}

func (u Anima) f2() int {
	//TODO implement me
	panic("implement me")
}

type User struct {
	name string
	age  int
}

func (u User) f2() int {
	panic("implement me")
}

// 非空接口 也可以表示任意类型
type IBase interface {
	f1() int
	f2() int
}

// 为结构体 User 构造方法
func (u User) f1() int {
	return 888
}

// 为结构体 User 构造方法
func (u Anima) f1() int {
	return 777
}

// 基于接口的参数，传入实际的结构体也必须有方法。 为什么不像 java 接口 实现呢。
func doingX(base IBase) {
	f1 := base.f1()
	fmt.Println(f1)
	f2 := base.f2() // User does not implement IBase (missing f2 method)
	fmt.Println(f2)
}

func main() {
	// 多种类型
	doingX(User{"flynn", 2})
	doingX(Anima{"flynn", 2})
}
