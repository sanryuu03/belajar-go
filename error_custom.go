package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "san" {
		return &notFoundError{"data not found"}
	}

	return nil // anggap saja ini ok
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// terjadi error
		if validationErr, ok := err.(*validationError); ok { // err coba di konversi ke pointer validationError
			// jika ok, maka validationErr ada isinya
			fmt.Println("validation error:", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Message)
		} else {
			fmt.Println("unknow error:", err.Error())
		}
	} else {
		// sukses
		fmt.Println("sukses")
	}

	kesalahan := SaveData("san", nil)
	if kesalahan != nil {
		switch finalError := kesalahan.(type) {
		case *validationError:
			fmt.Println("validation error with switch case:", finalError.Message)
		case *notFoundError:
			fmt.Println("not found error with switch case:", finalError.Message)
		default:
			fmt.Println("unknow error with switch case:", err.Error())

		}
	} else {
		// sukses
		fmt.Println("sukses")
	}
}
