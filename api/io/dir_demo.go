package main

import (
	"fmt"
	"os"
)

func main() {

	err := os.Mkdir("x2", 0755)
	fmt.Println(err)
	if err != nil {
		fmt.Printf("创建目录成功了")
		return
	} else {
		fmt.Printf("创建目录失败")
	}

	err2 := os.Remove("x2")
	if err2 != nil {
		fmt.Printf("目录删除成功了")
		return
	}

	// 多层
	//os.MkdirAll("t2/abc/dex/test.txt", 0755)
	//os.RemoveAll("t2/abc/dex/test.txt")

	fileInfo, err := os.Stat("x2")
	if err != nil {
		boolDir := fileInfo.IsDir()
		if boolDir {
			fmt.Printf("创建了是个目录")
		}
		return
	}
}
