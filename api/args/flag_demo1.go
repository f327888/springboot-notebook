package main

import (
	"flag"
	"fmt"
)

func main() {
	// demo -h -p -e
	host := flag.String("h", "127.0.0.1", "主机名")
	port := flag.Int("p", 8080, "端口")
	var email string
	flag.StringVar(&email, "e", "", "邮箱")
	flag.Parse()

	fmt.Println(*host, *port, email)
}
