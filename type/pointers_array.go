package main

import (
	"fmt"
	"reflect"
)

// 指针是一种数据类型
// 用于表示数据的内存地址

func main() {
	// 声明一个 字符串类型的变量并赋值
	var dataArray = [3]int8{11, 22, 33}

	fmt.Printf("%p \n", &dataArray)                                // 0xc0000aa058
	fmt.Printf("%p \n", &dataArray[0])                             // 0xc0000aa058 相等
	fmt.Printf("整个数组的类型 %s \n", reflect.TypeOf(&dataArray))        // *[3]int8
	fmt.Printf("数组的第一个元素的类型 %s \n", reflect.TypeOf(&dataArray[0])) // *[3]int8 地址相等，类型不一致

}
