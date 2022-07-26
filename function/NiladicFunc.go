package main

/*
 niladic 函数，翻译？？
函数定义时，它的形参一般是有名字的，不过我们也可以定义没有形参名的函数，只有相应的形参类型。
没有参数的函数通常被称为 niladic 函数 (niladic function)，就像 main.main() ？这个举例，没有形参呀
1. 意义、好处是什么呢？ 毕竟没有名字，参数也就利用不起来。和 Go 的简单的理念冲突。
*/

func f(int, int, string) int {
	return 100
}
func main() {
	f(1, 2, "name")
}
