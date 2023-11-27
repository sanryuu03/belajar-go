package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "san"
	middleName = "ryuu"
	lastName = "fullstack"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
