package main

import "fmt"

func ReadWrite() bool {
	file.Open("file.txt")
	// 做一些工作
	if failureX {
		file.Close()
		return false
	}
	if failureY {
		file.Close()
		return false
	}
	file.Close()
	return true
}
func main(){

}
