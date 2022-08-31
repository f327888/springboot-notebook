package main

func getFn() func(int, int) string {
	return exec
}
func exec(i int, i2 int) string {
	// ...
	return "ok"
}

// 函数类型
type f1 func(int, int) string

func getFn2() f1 {
	return fn1
}

func fn1(i int, i2 int) string {
	// ...
	return "ok"
}

func main() {
	getFn()
}
