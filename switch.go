package main

import "fmt"

func main() {
	name := "san"

	switch name {
	case "san":
		fmt.Println("hello san")
	case "ryuu":
		fmt.Println("hello ryuu")
	default:
		fmt.Println("hi, boleh kenalan?")
	}

	// switch short statement
	fmt.Println("switch short statement")

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	// switch tanpa kondisi
	fmt.Println("switch tanpa kondisi")
	job := "fullstack"
	lengthJob := len(job)
	switch {
	case lengthJob > 10:
		fmt.Println("nama terlalu panjang")
	case lengthJob > 5:
		fmt.Println("nama lumayan panjang")
	case false:
		fmt.Println("nama sudah benar")
	}
}
