package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blaclist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("san", blaclist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
