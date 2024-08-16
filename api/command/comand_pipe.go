package main

import (
	"fmt"
	"io"
	"os/exec"
)

/*
*
实时输出
*/
func main() {
	//cmdPtr := exec.Command("ls", "-s", "-l")
	cmd := exec.Command("ping", "-n", "3", "baidu.com")
	outPipe, _ := cmd.StdoutPipe()

	cmd.Start() // 和 Run 的区别

	for {
		chunk := make([]byte, 1024)
		_, err := outPipe.Read(chunk)
		if err != nil || err == io.EOF {
			// 读结束退出
			break
		}
		fmt.Print(string(chunk))
	}

	cmd.Wait()
}

//
//type Charset string
//
//const (
//	UTF8    = Charset("UTF-8")
//	GB18030 = Charset("GB18030")
//)
//
//func ConvertByte2String(byte []byte, charset Charset) string {
//
//	var str string
//	switch charset {
//	case GB18030:
//		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
//		str = string(decodeBytes)
//	case UTF8:
//		fallthrough
//	default:
//		str = string(byte)
//	}
//
//	return str
//}
