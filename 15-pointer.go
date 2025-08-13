package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main15() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // copy value [KARENA INI PASS BY VALUE]

	address2.City = "Bandung"
	fmt.Println(address2) // city berubah jadi bandung 
	fmt.Println(address1) // tidak berubah [KARENA INI PASS BY VALUE]

	// address1 akan berubah jika pakai POINTER bukan PASS BY VALUE
	// POINTER = kemampuan buat refrence ke lokasi memori yg sama tanpa duplikasi data [PASS BY REFERENCE]
	// cata buat POINTER pakai & sebelum reference nya
	address3 := &address1

	address3.City = "Pontianak"
	fmt.Println(address3) // city berubah jadi pontianak 
	fmt.Println(address1) // akan berubah jadi pontianak juga karena sudah di ubah sama address 3 di memori yg sama	

	// kalau POINTER mengacu ke struct baru, maka referencenya akan kembali ke semula
	address5 := &address1

	// kalau di assign ke struct baru
	address5 = &Address{"Jakarta","DKT Jakarta","Indonesia"}

	fmt.Println(address5) // city berubah jadi Jakarta 
	fmt.Println(address1) // akan masih tangerang karena di assign sama address3

	fmt.Println("-----------------------------------------------")

	// OPRATOR *, Akan mengubah seluruh value dari referencenya, INI BUKAN POINTER LAGI
	address6 := &address1
	*address6 = Address{"Papua","DKT Jakarta","Indonesia"}
	
	fmt.Println(address6) // papua
	fmt.Println(address1) // akan ikut jadi papua
	fmt.Println(address2) // bandung krna bkn pointer
	fmt.Println(address3) // akan ikut jadi papua karena pointer ke address1
	fmt.Println(address5) // jakarta krna pointer ke struct baru

	fmt.Println("-----------------------------------------------")

	// OPRATOR NEW 
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.Country = "TESSSSSS"

	fmt.Println(alamat1) // ini akan buat address jadi kosong
	fmt.Println(alamat2) // hasilnya akan sama karna pointer ke address1

	fmt.Println("-----------------------------------------------")

	// POINTER DI FUNCTION 
	alamatKu := Address{}
	ChangeCountryToIndonesiaValue(alamatKu)
	fmt.Println(alamatKu) // ga berubah, karena yang di masukkan ke func itu pass by value (duplikat dari value)

	alamatMu := &Address{}
	ChangeCountryToIndonesiaPointer(alamatMu)
	fmt.Println(alamatMu) // berubah karena yang dikirim itu pointer

	fmt.Println("-----------------------------------------------")

	// POINTER DI METHOD 
	eko := Man{"Eko"}
	eko.Married()
	fmt.Println(eko.Name) // akan tetap Eko, bukan Mr. Eko

	eko.MarriedPointer()
	fmt.Println(eko.Name) // akan tetap Eko, bukan Mr. Eko
}

// POINTER DI FUNCTION 
func ChangeCountryToIndonesiaValue(address Address){ // address itu pass by value
	address.Country = "Indonesia"
}

func ChangeCountryToIndonesiaPointer(address *Address){ // address ini pass by refrence, jadi akan ngubah value awalnya juga
	address.Country = "Indonesia"
}

// POINTER DI METHOD : NOTE. DIREKOMENDAISKAN SELALU PAKAI POINTER KALAU DI METHOD
type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr." + man.Name // ga ubah ke Mr beneran karena parameternya itu bukan pointer
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr." + man.Name // ga ubah ke Mr beneran karena parameternya itu bukan pointer
}
