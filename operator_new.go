package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := new(Address)
	address2 := address1

	address2.Country = "indonesia"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi indonesia
}
