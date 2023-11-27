package main

import "fmt"

func getFullName() (string, string) {
	return "san", "ryuu"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	fmt.Println("menghiraukan return value")
	nameDepan, _ := getFullName()
	fmt.Println(nameDepan)

	_, nameBelakang := getFullName()
	fmt.Println(nameBelakang)
}
