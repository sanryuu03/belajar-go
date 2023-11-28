package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // untuk tipe data interface, function, map, slice, pointer dan channel
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}

	data = NewMap("san")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
		fmt.Println(data["name"])
	}
}
