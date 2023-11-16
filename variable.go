package main

import "fmt"

func main() {
	// cara 1 deklarasi variable
	var name string

	name = "sanryuu"
	fmt.Println(name)

	name = "bersama selir membangun bangsa"
	fmt.Println(name)

	// cara 2 deklarasi variable
	var job = "backend"

	fmt.Println(job)

	job = "BE Golang"
	fmt.Println(job)

	// cara 3 deklarasi variable
	techstack := "laravel"

	fmt.Println(techstack)

	techstack = "laravel, express js"
	fmt.Println(techstack)

	// cara 4 deklarasi multiple variable
	var (
		fe = "react"
		be = "laravel"
	)

	fmt.Println(fe)
	fmt.Println(be)

	fe = "react js, next js"
	be = "ci, express js"
	fmt.Println(fe)
	fmt.Println(be)
}
