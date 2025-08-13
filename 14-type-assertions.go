package main

import "fmt"

func main14() {
	var result any = random()
	var resultString = result.(string)

	fmt.Println(resultString)
	
	// var resultInt = result.(int)
	// fmt.Println(resultInt) //bakal error panic karena aslinya itu tipe string, gabisa konversi ke int

	//kita bisa pakai switch untuk jaga" agar ga kena panic
	result2 := random()
	switch value := result2.(type){
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unkown", value)
	}
}

func random() any {
	return "OK"
}