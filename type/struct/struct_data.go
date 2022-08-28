package main

import "fmt"

// 1 定义
type Person struct {
	name  string
	age   int
	email string
	//addr  Address
}

type Address struct {
	city, state string
	code        int
}

func main() {
	// 2 初始化
	var p1 = Person{"冯大海", 32, "aaa@bb.com"}
	// 3 取值
	fmt.Println(p1.name, p1.age, p1.email)

}
