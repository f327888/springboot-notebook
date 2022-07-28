package main

import "fmt"

/*
值可以是任意类型的，这里给出了一个使用 func() int 作为值的 map：
*/
func main() {
	var mf = map[int]func() int{
		1: func() int {
			return 10
		},
		2: func() int {
			return 20
		},
		3: func() int {
			return 30
		}, // 居然不能少？ 和json 格式规范 不一致?
	}
	// 输出结果为：map[1:0x10903be0 5:0x10903ba0 2:0x10903bc0]: 整型都被映射到函数地址。
	fmt.Println(mf)
}
