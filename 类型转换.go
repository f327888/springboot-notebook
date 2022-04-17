package main

import (
	"fmt"
	"math"
)
/*  语法 : T(表达式)
 T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

 */
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	sqrtDemo()
}