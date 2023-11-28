package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"medan", "sumatera utara", "INA"}
	address2 := &address1 // pointer
	address2.City = "zogza"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi zogza

	fmt.Println("===========")
	fmt.Println("asterisk operator")
	fmt.Println("===========")

	*address2 = Address{"zogza", "DIY", "INA"}

	fmt.Println(address1) // address 1 berubah
	fmt.Println(address2)
}
