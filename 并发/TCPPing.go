package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var ip = "192.168.1.3"

func main() {
	//singleThreadPing()
	multiThreadPing()
}

/*
  都线程
*/
func multiThreadPing() {
	var begin = time.Now()
	var wg sync.WaitGroup

	for n := 21; n < 120; n++ {
		wg.Add(1) //添加wg
		go func(i int) {
			defer wg.Done() //释放wg
			var address = fmt.Sprintf("%s:%d", ip, i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println(address, " 打开")
		}(n)
	}
	wg.Wait() // 等待wg
	var elapseTime = time.Now().Sub(begin)
	fmt.Println("耗时", elapseTime)
}

/*
输出，单线程
192.168.1.3:23  是关闭的
192.168.1.3:24  是关闭的
192.168.1.3:25  是打开的
192.168.1.3:26  是关闭的

*/
func singleThreadPing() {
	// 本机ip

	for i := 21; i < 120; i++ {
		// ip:端口
		var address = fmt.Sprintf("%s:%d", ip, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Println(address, " 是关闭的")
			continue
		}
		conn.Close()
		fmt.Println(address, " 是打开的")
	}
}
