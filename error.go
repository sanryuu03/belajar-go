package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagian dengan nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := Pembagian(10, 5)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}
