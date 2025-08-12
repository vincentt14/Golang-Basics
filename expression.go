package main

import "fmt"

func main9(){
	// if expression
	name := "123456"

	if name == "da"{
		fmt.Println("lalalal")
	}else if name == "11" {
		fmt.Println("123213")
	}else {
		fmt.Println("dadadad")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	}else {
		fmt.Println("nama sudah benar")
	}

	// switch expression
	makanan := "nas"
	switch makanan{
	case "bakso": 
		fmt.Println("tong tong tong")
	case "nasgor": 
		fmt.Println("tek tek")
	default:
		fmt.Println("kring kring")
	}

	// switch short statement
	switch length := len(makanan); length > 5 {
	case true:
			fmt.Println("makannn")
	case false:
			fmt.Println("minumm")
	}

	gorengan2 := "geprek"
	length2 := len(gorengan2)
	switch {
	case length2 > 5:
		fmt.Println("goreng")
	case length2 > 3:
		fmt.Println("rebus")
	default:
		fmt.Println("ga makan")
	}

	// for loops
	counter := 1
	for counter <= 5 {
		fmt.Println("counter ke - ", counter)
		counter++
	}

	for counter2 := 20; counter2 >= 15; counter2-- {
		fmt.Println("angka ke - ", counter2)
	}

	// for range -> iterasi ke array / slice / map
	sliceNames := []string{"Eko","Kurniawan","Khannedy"}
	
	//manual
	for i := 0; i < len(sliceNames); i++ {
		fmt.Println(sliceNames[i])
	}

	//for range
	for index, name := range sliceNames {
		fmt.Println("index", index, "=", name)
	}

	//for range kalau ga butuh variabel ganti jadi _
	for _, name := range sliceNames {
		fmt.Println("name", name)
	}

	// break & continue
	for a:= 0; a < 10; a++ {
		if a % 2 == 0{
			continue
		}else if a == 9{
			break
		}
		fmt.Println("Perulangan ke", a)
	}
}