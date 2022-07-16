package main

import "fmt"

func intSum(x int, y int) int {
	return x + y
}

func sayHello() {
	fmt.Println("Hello 沙河")
}

// 多返回值
func manyReturn(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名
func manyReturnRename(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)
}
