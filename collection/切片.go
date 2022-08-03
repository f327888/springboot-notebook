package main

import "fmt"

/*
切片（Slice）是一个拥有相同类型元素的可变长度的一种结构体。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
*/
// var 数组变量名 [元素数量]T
// var 切片name []T  ,对比数组，不需要数量

func main() {
	main3()
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较(slice can only be compared to nil)
}

func main2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}

func main3() {
	a := make([]int, 2, 10)
	n, err := fmt.Println(a) //[0 0]
	fmt.Println(n, err)
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

	var aa = [2]string{"aa", "bb"}
	var bb = []string{"aa", "bb"}
	fmt.Printf(" aa type %T \n", aa)
	fmt.Printf(" bb type %T \n", bb)
}
