package main

import "fmt"

func main10() {
	// dever function = function yg kita jadwalkan utk dieksekusi setelah sebuah func lain selesai dieksekusi
	// dever func akan selalu di eksekusi walaupun terjadi error di function yg dieksekusinya
	runApplication()

	// panic function = function bwt hentiin program (force stop)
	// biasa dipakai kalau ada error crusial
	// saat program di berhentikan sama panic funtion, Defer func tetap akan jalan
	runApp(false)

	// recover function = function utk menangkap data panic
	// dengan recover, proses panic akan dihentikan. sehingga program ga akan force stop
	runAppRecover(true)
}

func runApplication(){
	defer logging()
	fmt.Println("Run application")
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Ups error")
	}
}

func endApp(){
	fmt.Println("End app")
}

func runAppRecover(error bool){
	defer endAppRecover()
	if error {
		panic("ERROR")
	}
}

func endAppRecover(){
	fmt.Println("program selesai")
	message := recover()
	fmt.Println("terjadi error", message)
}

