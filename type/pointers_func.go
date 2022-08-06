package main

import "fmt"

// https://www.youtube.com/watch?v=sTFJtxJXkaY

func main() {
	a := 4
	squareVal(a)
	squareAdd(&a) //TODO 没太明白, &a *int ？
}

func squareAdd(p *int) {
	*p = (*p) * (*p) // 中间有空格的 * 表示乘法
	fmt.Printf("i 的值: %v \n", *p)
}

func squareVal(v int) {
	v = v * v // 此处的 * 表示乘法
	fmt.Printf("i 的值: %v \n", v)
}
