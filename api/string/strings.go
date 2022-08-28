package main

import (
	"fmt"
	"strconv"
	"strings"
	myutf8 "unicode/utf8"
)

func main() {
	fields := strings.Fields("aa bb cc")
	fmt.Printf("--- %v \n", fields) //[aa bb cc]

	for index, value := range fields {
		fmt.Printf("index is %v, valuse is %v \n", index, value)
	}
	//index is 0, valuse is aa
	//index is 1, valuse is bb
	//index is 2, valuse is cc

	// 底层存储

	var name = "冯大海" // 一个汉字 3个字节

	// 获取字符串的长度  9(字节的长度) 一个字节 8位 bit
	fmt.Println(len(name))
	// 字符串转换为 字节的集合
	byteSet := []byte(name)
	fmt.Println(byteSet)

	//  字节的集合 转换为 字符串
	bytes := []byte{230, 173, 166, 230, 178, 155, 233, 189, 144}
	str := string(bytes)
	fmt.Println(str)
	// 汉字：冯
	//229 11100101
	//134 10000110
	//175 10101111
	fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	fmt.Println(name[1], strconv.FormatInt(int64(name[1]), 2))
	fmt.Println(name[2], strconv.FormatInt(int64(name[2]), 2))
	// 汉字：大
	//229 11100101
	//164 10100100
	//167 10100111
	fmt.Println(name[3], strconv.FormatInt(int64(name[3]), 2))
	fmt.Println(name[4], strconv.FormatInt(int64(name[4]), 2))
	fmt.Println(name[5], strconv.FormatInt(int64(name[5]), 2))
	// 汉字：海
	//230 11100110
	//181 10110101
	//183 10110111
	fmt.Println(name[6], strconv.FormatInt(int64(name[6]), 2))
	fmt.Println(name[7], strconv.FormatInt(int64(name[7]), 2))
	fmt.Println(name[8], strconv.FormatInt(int64(name[8]), 2))

	// rune 的集合 , 将字符串转换为 unicode 字符集编码点的集合，16进制
	// 显示是10进制， rune -- int32 , 4个字节
	tempSet := []rune(name)
	fmt.Println(tempSet)                                  //[20911 22823 28023]
	fmt.Println(strconv.FormatInt(int64(tempSet[0]), 16)) //51af -> 冯
	fmt.Println(strconv.FormatInt(int64(tempSet[1]), 16)) //5927 -> 大
	fmt.Println(strconv.FormatInt(int64(tempSet[2]), 16)) //6d77 -> 海

	// 反向 rune集合 转换成字符串
	runeList := []rune{20911, 22823, 28023}
	runeStr := string(runeList)
	fmt.Println(runeStr)
	// 字符串的长度,不是字节的长度
	runeLength := myutf8.RuneCountInString(name)
	fmt.Println(runeLength) // 3
}
