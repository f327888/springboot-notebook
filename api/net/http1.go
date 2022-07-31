package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
浏览器并输入 url 地址：http://localhost:8080/world，浏览器就会出现文字：Hello, world，
网页服务器会响应你在 :8080/ 后边输入的内容。
*/
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	//TODO F 是什么意思？
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:]) // 可以浏览器显示
}
func main() {

	http.HandleFunc("/", HelloServer) //TODO 函数的入参？？
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
