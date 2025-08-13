package main

import (
	"errors"
	"fmt"
)

func main18() {
	hasil, err := Pembagian(100, 0)
	if err == nil{
		fmt.Println("hasilnya adalah", hasil)
	}else {
		fmt.Println("ERROR: ", err.Error())
	}
}

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	}else {
		return nilai / pembagi, nil
	}
}