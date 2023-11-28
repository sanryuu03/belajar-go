package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var san Customer
	fmt.Println(san)

	san.Name = "san"
	san.Address = "medan"
	san.Age = 25

	fmt.Println(san)
	fmt.Println(san.Name)
	fmt.Println(san.Address)
	fmt.Println(san.Age)

	fmt.Println("struct literals")

	ryuu := Customer{
		Name:    "ryuu",
		Address: "zogza",
		Age:     25,
	}
	fmt.Println(ryuu)

	sang := Customer{"sang wolves", "INA", 30}
	fmt.Println(sang)
}
