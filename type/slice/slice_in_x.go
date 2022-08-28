package main

import "fmt"

func main() {

	v1 := []int{11, 22, 33}

	v2 := [][]int{[]int{1, 2, 3}, []int{77, 88, 99}}
	v3 := [][2]int{[2]int{1, 2}, [2]int{77, 88}}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	v1[0] = 1111
	v2[0][2] = 22222
	v3[1][0] = 9999

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

}
