package main

import "fmt"

func endApp() {
	fmt.Println("end app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ups error")
	}
}

func main() {
	runApp(true)
}
