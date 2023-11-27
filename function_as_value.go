package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("san"))
}
