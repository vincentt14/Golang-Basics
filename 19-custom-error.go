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

func main() {
	err := SaveData("", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Error())
		}else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Error())
		}else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("Sukses")
	}

}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "Eko" {
		return &notFoundError{"Data not found"}
	}

	return nil
}