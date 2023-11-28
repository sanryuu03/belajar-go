package main

import (
	"belajar-go/helper"
	"fmt"
)

/**
 * di go-lang, untuk menentukan access modifier, cukup dengan nama function atau variable
 * jika namanya di awali dengan huruf besar, maka artinya bisa di akses dari package lain,
 * jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain
 */

func main() {
	result := helper.SayHai("san")
	fmt.Println(result)
	fmt.Println(helper.Application)
	// fmt.Println(helper.version)           // tidak bisa di akses
	// fmt.Println(helper.sayGoodBye("san")) // tidak bisa di akses
}
