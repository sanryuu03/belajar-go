package main

import "fmt"

func main() {
	// cara 1 deklarasi variable
	var names [3]string
	names[0] = "san"
	names[1] = "ryuu"
	names[2] = "fullstack"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	// cara 2 deklarasi variable
	var values [3]int = [3]int{
		90,
		80,
		85,
	}

	fmt.Println(values)

	// cara 3 deklarasi variable
	var salary = [3]int{
		9000000,
		8000000,
		8500000,
	}

	fmt.Println(salary)

	// function array
	fmt.Println(len(salary))
	salary[2] = 9500000
	fmt.Println(salary)

	// cara 4 deklarasi variable => [...] => jumlahnya dinamis
	var ages = [...]int{
		17,
		19,
		20,
		25,
	}

	fmt.Println(ages)
	fmt.Println(len(ages))
}
