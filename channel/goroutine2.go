package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	//time.Sleep(2e9) // 打开 无输出
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
