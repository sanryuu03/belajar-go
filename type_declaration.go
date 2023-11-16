package main

import "fmt"

func main() {
	type NoKTP string

	var ktpSan NoKTP = "1234567890"

	fmt.Println(ktpSan)

	var dummy string = "88888888"
	var dummyKtp NoKTP = NoKTP(dummy)

	fmt.Println(dummyKtp)
}
