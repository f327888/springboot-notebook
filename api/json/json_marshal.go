package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	v1 := []interface{}{
		"flynn", 23, true, 99999.00, // 不是 key value ？？？ 和json 如何兼容？
		map[string]interface{}{
			"name": "阿海",
			"age":  23,
		},
		Person{Name: "hai", Age: 28},
	}
	marshal, _ := json.Marshal(v1)
	// 字节
	fmt.Println(marshal)
	data := string(marshal)
	fmt.Println(data) // ["flynn",23,true,99999,{"age":23,"name":"阿海"}]
	// Person 字段小写 则无法读取到， ["flynn",23,true,99999,{"age":23,"name":"阿海"},{}]
	//["flynn",23,true,99999,{"age":23,"name":"阿海"},{"Name":"hai","Age":28}]
	// 属性的字段也小写, 字段后面 `json:"name"`
}
