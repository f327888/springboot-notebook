package main

import "fmt"

func main() {
	name := [2]string{"冯大海", "flynn"}

	data := name[0]
	fmt.Println(data)
	data = "alex"
	fmt.Println(name)
	// 修改值
	name[0] = "alex"
	fmt.Println(name)

	// 切片
	nums := []int{123, 456, 789}
	ints := nums[0:2]
	fmt.Println(nums) // 不改变
	fmt.Println(ints) // [123 456]

}
