package main

import "fmt"

/*
   为什么要分3种写法，和java 一样，一种不是更简单吗
 */
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func main() {
	forDemo3()
}