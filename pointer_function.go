package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(alamat Address) {
	alamat.Country = "indonesia"
}

func ChangeCountryToIndonesiaPointer(alamat *Address) {
	alamat.Country = "indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndonesia(address)

	fmt.Println(address) // tidak berubah

	ChangeCountryToIndonesiaPointer(&address)
	fmt.Println(address) // berubah

	alamat := &Address{}
	ChangeCountryToIndonesiaPointer(alamat)

	fmt.Println(alamat) // berubah
}
