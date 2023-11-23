package main

import "fmt"

func main() {
	// cara 1 deklarasi variable
	var person1 map[string]string = map[string]string{}
	person1["name"] = "san"
	person1["job"] = "fullstack"

	// cara 2 deklarasi variable
	var person2 map[string]string = map[string]string{
		"name": "san",
		"job":  "react js",
	}

	// cara 3 deklarasi variable
	person3 := map[string]string{
		"name": "san",
		"job":  "go",
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Println(person3["job"])

	fmt.Println("========")
	fmt.Println("make map (buat map baru)")

	book := make(map[string]string)
	book["title"] = "fundamental golang"
	book["author"] = "sanryuu"
	book["wrong"] = "salah"

	fmt.Println(book)
	delete(book, "wrong")

	fmt.Println(book)
}
