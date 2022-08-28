package main

import (
	"fmt"
	"reflect"
)

// 1 定义
type person2 struct {
	name  string "姓名"
	age   int    "年龄"
	email string "邮件"
}

func main() {
	// 2 初始化
	var p1 = person2{"冯大海", 32, "aaa@bb.com"}
	p1Type := reflect.TypeOf(p1)
	//1 按索引
	field1 := p1Type.Field(0)
	fmt.Println(field1.Tag)
	//2 按名称
	field2, _ := p1Type.FieldByName("name")
	fmt.Println(field2.Tag)

}
