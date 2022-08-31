package main

import "fmt"

/**
匿名函数
*/

func fn11(n1 int, n2 int) func(int2 int) string {
	return func(int2 int) string {
		// ...
		return "匿名函数"
	}
}

func main() {

	vf := func(n1 int, n2 int) int {
		return n1 + n2
	}

	i := vf(11, 22)
	fmt.Println(i)

	v2 := func(n1 int, n2 int) int {
		return n1 + n2
	}(22, 33) // (直接执行)
	fmt.Println(v2)

	// 返回一个函数
	f3 := fn11(1, 2)
	s := f3(1)
	fmt.Println(s)
}
