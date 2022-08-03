package main

import (
	"fmt"
	"strings"
)

func main() {
	fields := strings.Fields("aa bb cc")
	fmt.Printf("--- %v \n", fields) //[aa bb cc]

	for index, value := range fields {
		fmt.Printf("index is %v, valuse is %v \n", index, value)
	}
	//index is 0, valuse is aa
	//index is 1, valuse is bb
	//index is 2, valuse is cc

}
