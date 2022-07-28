package main

import "fmt"

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
		return
	}
	if score > 75 {
		fmt.Println("B")
		return
	}
	fmt.Println("C")
}

// 和 上面的区别是什么？

func ifDemo2() {
	//if 表达式之前添加一个执行语句，再根据变量值进行判断
	if score := 115; score >= 190 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func main() {
	ifDemo1()
	ifDemo2()
}
