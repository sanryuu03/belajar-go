package main

import "fmt"

func main() {
	names := [...]string{"san", "ryuu", "fe", "be", "fullstack", "golang"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	var slice5 []string = names[:]
	fmt.Println(slice5)

	//function slice
	fmt.Println(len(slice1))
	fmt.Println(len(slice5))

	fmt.Println("========")

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println("daySlice1", daySlice1)

	fmt.Println(days)

	fmt.Println("========")

	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println("daySlice2", daySlice2)
	daySlice2[0] = "ups"
	fmt.Println("daySlice2", daySlice2)
	fmt.Println(days)

	fmt.Println("========")
	fmt.Println("make slice (buat slice dari awal)")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "san"
	newSlice[1] = "ryuu"
	// newSlice[2] = "be" => error karna panjangnya sudah "2", harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "be")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "hasan"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("========")
	fmt.Println("copy slice")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fmt.Println("========")
	fmt.Println("array vs slice")

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
