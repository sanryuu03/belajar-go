package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAllSlice(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAllSlice([]int{10, 10, 10}))

	fmt.Println("slice parameter di variadic function")
	numbers := []int{10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
