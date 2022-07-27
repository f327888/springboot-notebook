package main

import "fmt"

func intSum(x int, y int) int {
	return x + y
}

//返回值命名 需要加 （）
func intSum2(x int, y int) (aa int) {
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
