package main

import (
	_ "belajar-go/blank_identifier"
	"belajar-go/database"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
