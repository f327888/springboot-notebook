package main

import "fmt"

func main() {
	v1 := make([]int, 1, 3)
	fmt.Println(len(v1), cap(v1)) //1 3
	// 长度最大是2 ，索引是长度范围的索引，即实际元素的索引
	v2 := make([]int, 2, 5)
	fmt.Println(v2[0]) //1 3
	fmt.Println(v2[1]) //1 3
	//fmt.Println(v2[2]) // 报错

	// 切片的切片, 存储数据的地方是一致的
	v11 := []int{11, 22, 33, 44, 55, 66}
	v22 := v11[2:3]
	v33 := v11[1:]
	v44 := v11[:3]
	fmt.Println(v22) // [33]
	fmt.Println(v33) // [22 33 44 55 66]
	fmt.Println(v44) // [11 22 33]

	v11[2] = 999
	fmt.Println(v11) //
	fmt.Println(v22) //
}
