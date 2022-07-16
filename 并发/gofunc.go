package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	say("hello")
	go say("world") // 没有 go , 先输出3个 hello,接着再输出 world
	//time.Sleep(5000 * time.Millisecond) // 输出 world， 注释掉则不输出。主进程直接关闭, 对应goroutine也直接关闭,导致函数没有运行
	fmt.Println("------执行结束------")
}
