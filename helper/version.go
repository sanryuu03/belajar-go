package helper

var version = "1.0.0" // tidak bisa di akses dari luar
var Application = "golang"

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "good bye " + name
}

// bisa diakses dari luar package
func SayHai(name string) string {
	return "hello " + name
}
