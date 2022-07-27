package main

import (
	"fmt"
	"time"
)

// 利用缓存，数据结构 数组的 index 是连续的 n ,和 斐波那契数一致。
/*
 第一次是否利用到了缓存： 否
第二次和第一次有区别吗：时间差：4.7349ms
*/
const LIM = 41

// 数组 用于缓存结果 index(n)-->result
var fibs [LIM]uint64

func main() {
	funcWithCache()
	//funcName()
}

/*
 不缓存：2.3832401s
  缓存时间怎么还变长了？数值小体现不出来？
但是 教程里面是：
Output: LIM=40:
normal (fibonacci.go): the calculation took this amount of time: 4.730270 s
     with memoization: the calculation took this amount of time: 0.001000 s
*/
func funcWithCache() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci2(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res // 放到数组缓存中
	return
}
