package main

import "fmt"

func Ups() interface{} {
	// return 1
	// return true
	return "ups"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
