package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	// augmented assignments
	var i = 10
	i += 10
	fmt.Println(i)

	// unary operator
	var j = 1
	fmt.Println(j)
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}
