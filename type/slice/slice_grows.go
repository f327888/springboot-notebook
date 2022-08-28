package main

import "fmt"

func main() {
	v1 := make([]int, 1, 3)
	v2 := append(v1, 66)

	fmt.Println(v1) // [0]
	fmt.Println(v2) // [0 66]
	v1[0] = 999
	fmt.Println(v1) // [999]
	fmt.Println(v2) // [999 66]

	// 需求：往一个切片追加一个
	v3 := make([]int, 1, 3)
	v3 = append(v3, 66)
	fmt.Println(v3) // [999]

	///////////////////////////
	v11 := []int{11, 22, 33} // [3]长度 错误？？
	v22 := append(v11, 44)
	fmt.Println(v22) //
}
