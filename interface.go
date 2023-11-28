package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (orang Person) GetName() string {
	return orang.Name
}

type Animal struct {
	Name string
}

func (hewan Animal) GetName() string {
	return hewan.Name
}

func main() {
	person := Person{Name: "san"}
	SayHello(person)

	animal := Animal{"kucing"}
	SayHello(animal)
}
