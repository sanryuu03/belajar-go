package main

import "fmt"

func main() {
	const name string = "san"
	const job = "be"

	fmt.Println(name)
	fmt.Println(job)

	const (
		fe string = "react"
		be        = "laravel"
	)

	fmt.Println(fe)
	fmt.Println(be)
}
