package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 /Go"
	format                 = "%f %d % s"
)

func main() {
	fmt.Println("Please enter your full name:")
	fmt.Scanln(&firstName, &lastName)              //Scanln() 扫描来自标准输入的文本 为什么要用 &
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
}
