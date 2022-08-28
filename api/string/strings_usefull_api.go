package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 长度：字节和字符的长度
	name := "冯大海"
	fmt.Println(len(name))                    // 9
	fmt.Println(utf8.RuneCountInString(name)) // 3

	nickname := "fly nn"
	fmt.Println(strings.ToLower(nickname))
	fmt.Println(strings.ToUpper(nickname))
	fmt.Println(strings.Trim(nickname, "f"))
	// trim 截取
	fmt.Println(strings.TrimPrefix(nickname, "f"))
	fmt.Println(strings.TrimLeft(nickname, "f"))
	fmt.Println(strings.TrimRight(nickname, "f"))
	fmt.Println(strings.TrimSpace(nickname))
	//strings.TrimFunc(nickname, myTrim(rune(name)))
	// Replace 替换
	// Split 分割
	// Join
	// builder 性能好
}

func myTrim(r rune) bool {

	return true
}
