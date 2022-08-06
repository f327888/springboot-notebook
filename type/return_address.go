package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() person {
	m := person{
		name: "noname",
		age:  50,
	}
	return m // 返回值
}

func initPerson2() *person {
	m := person{
		name: "noname",
		age:  50,
	}
	return &m // 返回地址  & 和 * 自动转换？
}

func main() {
	fmt.Println(initPerson())  // {noname 50}
	fmt.Println(initPerson2()) // &{noname 50}
}
