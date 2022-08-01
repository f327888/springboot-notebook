package main

import "fmt"

var aa bool = false

func main() {

	//if true {
	//	aa := true // aa declared but not used  新的变量 aa 和外面的全局变量重名，所以错误
	//}
	//fmt.Println(aa)

	if true {
		aa = true // aa declared but not used
	}

	fmt.Println(aa)
}
