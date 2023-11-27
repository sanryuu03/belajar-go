package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	fmt.Println("selesai")

	// for dengan statement
	fmt.Println("for dengan statement")

	for hitung := 1; hitung <= 10; hitung++ {
		fmt.Println("perulangan hitung ke ", hitung)
	}

	fmt.Println("selesai hitung")

	// for range array, slice, map
	fmt.Println("for range array, slice, map")

	names := []string{"san", "ryuu", "full stack"}

	// for manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// for range tanpa index
	for _, nama := range names {
		fmt.Println(nama)
	}
}
