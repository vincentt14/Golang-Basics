package main

import "fmt"

type HasName interface {
	GetName() string
}

// ini termasuk interface udah memenuhi kontrak
func SayHello(value HasName) {
	fmt.Println("hello", value.GetName())
}

// cara implementasi interface itu, harus buat struct yang punya method di interface (sesuai kontrak)
type Person struct {
	Name string
}

// contoh implementasi di Person (BUAT FUNC yang ada di kontrak)
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main12() {
	// interface = tipe abstrak yg berisi definisi method (func di struct)
	// interface itu bersifat kontrak (implementasi struct nya bole banyak yg penting ngikut kontrak)
	person := Person{Name: "Cent"}
	SayHello(person)

	animal := Animal{Name: "kucing"}
	SayHello(animal)

	// interface kosong
	var kosong = Ups()
	fmt.Println(kosong)	
}

// interface kosong
func Ups() any {
	return 1
}