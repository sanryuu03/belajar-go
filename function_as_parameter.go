package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

type Filter func(string) string

func sayHelloWithFilterTypeDeclarations(nama string, filter Filter) {
	fmt.Println("hello", filter(nama))
}

func main() {
	sayHelloWithFilter("san", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
	fmt.Println("function type declarations")
	sayHelloWithFilterTypeDeclarations("ryuu", filter)
}
