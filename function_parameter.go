package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

func main() {
	sayHelloTo("san", "ryuu")
}
