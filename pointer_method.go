package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	san := Man{"san"}
	san.Married()
	fmt.Println(san)

	san.MarriedPointer()
	fmt.Println(san)
}
