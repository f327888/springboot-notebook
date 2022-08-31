package main

import "fmt"

/**
闭包
*/

func main() {
	var fnList []func()
	// 只是创建函数，函数里面没执行。
	for i := 0; i < 5; i++ {
		fn := func() {
			fmt.Println(i)
		}
		fnList = append(fnList, fn)
	} // i = 5
	// 运行 采取执行具体的函数
	fnList[0]() // 5
	fnList[1]() // 5
	fnList[2]() // 5

	//// 想不相互影响 0 1 2 。。。 需要用 闭包,先存起来，
	//var fnList2 []func()
	//// 只是创建函数，函数里面没执行。
	//for i := 0; i < 5; i++ {
	//	fn := func(arg int) {
	//	}(i) // 执行了 第一次 arg =0
	//	fnList2 = append(fnList2, fn)
	//} // i = 5
	//// 运行 采取执行具体的函数
	//fnList2[0]() //
	//fnList2[1]() //
	//fnList2[2]() //

	var fnList3 []func()
	// 只是创建函数，函数里面没执行。
	for i := 0; i < 5; i++ {
		fn := func(arg int) func() {
			return func() {
				fmt.Println(arg)
			}
		}(i) // 执行了 i 作为匿名函数的参数，即 i->arg
		fnList3 = append(fnList3, fn)
	} // i = 5
	// 运行 采取执行具体的函数
	fnList3[0]() // 0
	fnList3[1]() // 1
	fnList3[2]() // 2

}
