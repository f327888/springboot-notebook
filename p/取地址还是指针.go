package main

import "fmt"

/*
	&和*区别，比如 p , *P , &p
*/

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) size() float64 {
	return r.width * r.height
}

func main() {
	// &是取地址符号, 取到Rect类型对象的地址
	fmt.Println(&Rect{100, 50}) // &{100 50}  有 &
	fmt.Println(Rect{100, 50})  // {100 50}   无 &
	//fmt.Println(*Rect{100, 50})   // 编译报错
	fmt.Println(*(&Rect{100, 50})) // {100 50} 无 &

	// *可以表示一个变量是指针类型(r是一个指针变量):
	var r *Rect = &Rect{100, 50}
	var r2 = &Rect{100, 50} // 和上面等价

	fmt.Println(r)        // &{100 50}
	fmt.Printf("%p\n", r) // 0xc0000160f0
	fmt.Println(r2)       // &{100 50}

	// *也可以表示指针类型变量所指向的存储单元 ,也就是这个地址所指向的值
	fmt.Println(*r)  // {100 50}
	fmt.Println(*r2) // {100 50}

	// 查看这个指针变量的地址 , 基本数据类型直接打印地址
	fmt.Println(&r) // 0xc00012c020  0x 16进制格式的地址

}
