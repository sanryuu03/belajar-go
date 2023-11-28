package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"medan", "sumatera utara", "INA"}
	address2 := address1 // copy value

	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("===========")

	address2.City = "zogza"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah menjadi zogza

	fmt.Println("===========")
	fmt.Println("pointer")
	fmt.Println("===========")

	var address3 Address = Address{"medan", "sumatera utara", "INA"}
	var address4 *Address = &address3 // pointer

	fmt.Println(address3)
	fmt.Println(address4)

	fmt.Println("===========")

	address4.City = "zogza"
	fmt.Println(address3) // ikut berubah
	fmt.Println(address4) // berubah jadi zogza
}
