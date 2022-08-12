package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 0
	var size = unsafe.Sizeof(i)
	fmt.Println("The size of an int is: ", size)
	alignof := unsafe.Alignof(i)
	fmt.Println("The alignof of an int is: ", alignof)
	//fmt.Println("The alignof of an int is: ", unsafe.Pointer)
	//fmt.Println("The alignof of an int is: ", unsafe.IntegerType)
}
