package main

func main() {
	ch1 := make(chan string)
	// int类型的通道的通道
	chch1 := make(chan chan int)

	// 函数通道
	funcChan := make(chan func())

}
