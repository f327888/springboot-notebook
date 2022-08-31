package main

import "fmt"

// 结构体 + 函数： 通过函数创建结构体
// 结构体工厂
type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	// ... 结构体很多逻辑
	return &File{fd, name} // & 地址
}

func main() {

	file1 := NewFile(10, "./test1.txt")
	file2 := File{10, "./test1.txt"}
	// 问题？都可以，风格不统一 --> 对象改为私有的。
	fmt.Println(file1)
	fmt.Println(file2)
}
