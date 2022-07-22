package main

import "strconv"

func getInt(s string, _default int) int {
	if s == "" {
		return _default
	}
	int, err := strconv.Atoi(s) // int 可以作为变量名
	if err != nil {
		//log.Errorf("Invalid int value: %s", err)
		return _default
	}
	return int
}

func main() {

}
