package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	json_str := "[\"flynn\",23,true,99999,{\"age\":23,\"name\":\"阿海\"},{\"name\":\"hai\",\"age\":28}]\n"
	// [] 切片或者 数组;  大小不确定 -- 切片
	// 不是具体的结构体, 里面是 map这种通用的结构体
	var json_object []interface{} // [flynn 23 true 99999 map[age:23 name:阿海] map[age:28 name:hai]]
	//var json_object map[string]interface{} // map[]

	json.Unmarshal([]byte(json_str), &json_object)
	fmt.Println(json_object)

}
