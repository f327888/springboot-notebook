package main

import "fmt"

/**
map 是引用类型，可以使用如下声明：
var map1 map[keytype]valuetype
var map1 map[string]int
初始化：


key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float32(64)。
所以数组、切片和结构体不能作为 key (译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），
但是指针和接口类型可以。
如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。
*/
func main() {
	// 先声明再赋值
	var nameScoreMap map[string]int
	nameScoreMap = map[string]int{"zhangsan": 90, "lisi": 85}
	fmt.Println(nameScoreMap["zhangsan"])
	fmt.Println("--------------------------------")
	//
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	// 如果key1存在则ok == true，否则ok为false(_ 忽略值)
	if _, ok1 := scoreMap["张三"]; ok1 {

	}
}
