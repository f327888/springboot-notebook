package main

// 用于 slice 切片，类似 java 的 arrayList
func main() {
	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1) // 新的切片，长度+1=32，容量（*2）=64
	a[2] = 42
}
