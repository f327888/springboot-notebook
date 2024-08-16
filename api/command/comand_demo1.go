package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	var out bytes.Buffer

	//cmdPtr := exec.Command("ls", "-s", "-l")
	cmdPtr := exec.Command("go", "version") // go version go1.18 windows/amd64
	// 目录切换
	cmdPtr.Dir = "/user/local/go/bing/go"
	cmdPtr.Stdout = &out
	cmdPtr.Run()

	result := out.String()
	fmt.Println(result)
}
