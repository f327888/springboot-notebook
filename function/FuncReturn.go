package main

/*
任何一个有返回值（单个或多个）的函数都必须以 return 或 panic（参考 第 13 章）结尾
*/

func fWithReturn(age int) int {
	return age + 1
}

type Stack []*int

func (st *Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
		return v
	}

	return v
}

func main() {

}
