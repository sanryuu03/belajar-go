package main

import "fmt"

func random() interface{} {
	return "ok"
}

func randomNumber() any {
	return 25
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	resultNumber := randomNumber()
	// ambil type
	switch value := resultNumber.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknow", value)
	}
}
