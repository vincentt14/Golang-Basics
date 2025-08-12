package main

import "fmt"

func main4(){
	var name string

	name = "Vincent"
	name = "Tes"
	fmt.Println(name)

	var iniNumber = 5
	fmt.Println(iniNumber)

	// simplify var declaration 
	deklarasiPertama := "halo"
	deklarasiPertama = "udah gabisa pakai := karena bukan deklarasi awal"
	fmt.Println(deklarasiPertama)

	// multiple variabel declaration
	var (
		firstName = "Eko"
		middleName = "Kurniawan"
		lastName = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	// const
	const makanan = "indomie"
	
	const (
		makanan1 = "sedap"
		makanan2 = "bubur"
	)
}