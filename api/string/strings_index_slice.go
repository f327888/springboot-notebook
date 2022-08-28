package main

import "fmt"

func main() {
	// 长度：字节和字符的长度
	name := "冯大海"
	// 索引获取字节 注意不要数组越界
	v1 := name[0]
	fmt.Println(v1)

	// 切片获取字节区间
	v2 := name[0:3]
	fmt.Println(v2) // 冯 3是长度？

	v3 := name[0:2]
	fmt.Println(v3) // 乱码

	// 获取所有字节  不太用的到 --> rune
	for i := 0; i < len(name); i++ {
		//0 229
		//1 134
		//2 175
		//3 229
		//4 164
		//5 167
		//6 230
		//7 181
		//8 183
		fmt.Println(i, name[i])
	}

	// 获取所有字符
	for index, item := range name {
		//0 20911 冯
		//3 22823 大
		//6 28023 海
		fmt.Println(index, item, string(item))
	}

	// 转换为 rune 集合
	runeList := []rune(name)
	fmt.Println(runeList[0], string(runeList[0]))
	fmt.Println(runeList[1], string(runeList[1]))
}
