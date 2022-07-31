package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "aaa.text"
	outputFile := "aaa_copy.text"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
		return
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
