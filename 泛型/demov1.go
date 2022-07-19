package main

import "fmt"

/*
输出数组的值，如果没有泛型的话，这个函数需要写出 int 版，float版，string 版，以及我们的自定义类型（struct）的版本
 @see https://coolshell.cn/articles/21615.html
*/
func print[T any](arr []T) {
	// range 迭代 xxx
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	strs := []string{"Hello", "World", "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718}
	nums := []int{2, 4, 6, 8}

	print(strs)
	print(decs)
	print(nums)
}
