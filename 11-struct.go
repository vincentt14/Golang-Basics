package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main11() {
	// buat data struct biasa
	var pelanggan Customer
	pelanggan.Name = "Vincent"
	pelanggan.Address = "Jalan A"
	pelanggan.Age = 24

	// buat data struct literal
	pelanggan2 := Customer {
		Name : "Cent",
		Address: "Indonesia",
		Age: 20,
	}

	pelanggan3 := Customer{"Budi", "Indonesia", 30}

	fmt.Println(pelanggan)
	fmt.Println(pelanggan2)
	fmt.Println(pelanggan3)

	pelanggan3.sayHello("Agus")
}

// function dalam struct => metod yg menempel di struct
func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}