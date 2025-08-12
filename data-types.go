package main

import "fmt"

func main8(){
	// 1. number
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)

	// 2. boolean
	fmt.Println("benar = ", true)
	fmt.Println("salah = ", false)

	// 3. string
	fmt.Println("ini string")
	fmt.Println(len("ini string"))
	fmt.Println("baba"[0])
	fmt.Println("baba"[1])

	// oprasi boolean
	var i = "lala"
	var apapun string = "makanan"

	fmt.Println(i != apapun) 

	// 4. array
	var names[3]string
	names[0] = "haha"
	names[1] = "hoho"
	names[2] = "hehe"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		30,
		//kalau kosong index ke 3 maka akan di isi default int = 0
	}

	fmt.Println(values)
	fmt.Println(values[0])

	var unknown = [...]int{
		100,
		20,
		50,
		22,
		70,
	}

	fmt.Println(unknown)
	fmt.Println(len(unknown))

	// 5. slice
	// array tapi ukurannya bisa bertambah, berhubungan karna slash = data yg mengakses ebagian / sluruh data di array
	// ada 3 poin utama = pointer, length, capacity
	names2 := [...]string{"Eko", "Kurniawan", "Khannedy", "Joko", "Budi", "Nugraha"}
	fmt.Println(names2)
	
	slice1 := names2[4:6]
	fmt.Println(slice1)

	slice2 := names2[:3]
	fmt.Println(slice2)

	var slice3 []string = names2[3:]
	fmt.Println(slice3)
	
	var slice4 []string = names2[:]
	fmt.Println(slice4)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	days := [...]string{"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Budi")
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(fromSlice, toSlice)

	fmt.Println(days)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Map. untuk akses slice / array. pairing key - value. Unique. kalau key sama maka akan ditimpa
	person := map[string]string{
		"name": "Vincent",
		"address" : "Ptk",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(person["dsadsads"]) //karna ga ada, jadi dia akan default value tipe nya, ""
	
	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "author")

	fmt.Println(book)
	fmt.Println(len(book))
}