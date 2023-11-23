package main

import "fmt"

func main() {
	name := "san"

	if name == "san" {
		fmt.Println("benar, hello san")
	} else if name == "ryuu" {
		fmt.Println("benar, hello ryuu")
	} else {
		fmt.Println("hi, boleh kenalan ?")
	}

	// if short statement
	fmt.Println("if short statement")

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
