package main

import (
	"fmt"
	"math"
)

/*
变量除了可以在全局声明中初始化，也可以在 init() 函数中初始化。
这是一类非常特殊的函数:
1. 它不能够被人为调用，
2. 而是在每个包完成初始化后自动执行，
3. 并且执行优先级比 main() 函数高。
每个源文件可以包含多个 init() 函数，
同一个源文件中的 init() 函数会按照从上到下的顺序执行，
如果一个包有多个源文件包含 init() 函数的话，则官方鼓励但不保证以文件名的顺序调用。
初始化总是以单线程并且按照包的依赖关系顺序执行。

一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。
*/
var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
	fmt.Println("...init method")
}

//...init method
//...main method
func main() {
	fmt.Println("...main method")
}
