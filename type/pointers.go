package main

import "fmt"

// https://www.youtube.com/watch?v=sTFJtxJXkaY

func main() {
	i := 42
	//fmt.Println(*i) // 错误
	fmt.Printf("i 的值: %v \n", i)
	fmt.Printf("i 的地址: %v \n", &i)
	p := &i // 声明一个变量 指向 i 的地址, 由于
	//p := i // *p 报错

	fmt.Printf("p 的值: %v \n", p)     // 即 &i, 地址 0xc000126058
	fmt.Printf("p 的取消引用: %v \n", *p) // *(&i) , 取消引用，
	fmt.Printf("p 的地址: %v \n", &p)   // p 自己的地址 0xc000006030

	//q := i
	//fmt.Println(*q) // 错误
	//w := *i         // 错误
	//
	//fmt.Println(q)
	//fmt.Println(&q)

}
