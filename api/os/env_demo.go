package main

import (
	"fmt"
	"os"
)

func main() {
	str := "My name is $Name and live in $City."
	strResult := os.ExpandEnv(str) // 读取 $var  ${var} 变量
	// 配置环境变量 Name=flynn;City=nanjing , 输出 My name is flynn and live in nanjing.
	fmt.Println(strResult)
}
