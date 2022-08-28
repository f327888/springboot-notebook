package main

import "fmt"

// https://www.youtube.com/watch?v=sTFJtxJXkaY

func main() {
	name := "flynn"
	// copy name 的值，赋值给 data
	changeData1(name)
	fmt.Printf("name 的值: %v \n", name)
	fmt.Printf("name 的值是否改变: %v \n", name == "flynn")

	name2 := "flynn"
	// copy name 的值，赋值给 data
	changeData2(&name2)
	fmt.Printf("name2 的值: %v \n", name2)
	fmt.Printf("name2 的值是否改变: %v \n", name2 == "flynn")
}
func changeData1(data string) {
	data = "大海" //
	//&data = "大海" // 错误，类型不一致 string <-> 0x.... 地址
	//*data = "大海" // 错误，类型不一致 string <-> *string
}

func changeData2(data *string) {
	//data = "大海" // 错误，类型不一致 string <-> *string
	//&data = "大海" // 错误，类型不一致 string <-> 0x.... 地址
	*data = "大海"
}
