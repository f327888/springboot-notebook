package main

import (
	. "fmt"
)

/*
   耗子大佬 https://coolshell.cn/articles/21146.html

*/

type User struct {
	Id       string // 必须
	Name     string // 必须
	Age      int    // 非必须
	Gender   bool   // 非必须
	password string // 小写，非导出
}

func main() {
	//for _, value := range os.Args {
	//	Println("Arg=", value)
	//}

	u := User{
		Id: "123", Name: "zhangsan", Age: 19, Gender: false, password: "123456",
	}
	Println(u)
}
