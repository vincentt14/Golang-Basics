package helper //nama package mengikuti nama folder

// Huruf besar = bisa di akses dari package lain
// Huruf kecil = tidak bisa di akses dari package lain

var Application = "Golang" // bisa diakses dari luar folder helper
var version = "1.0.0" // gabisa karna hruf kecil

func SayHello(name string) string{ // bisa diakses dari luar folder helper
	return "hello " + name
}

func sayGoodbye(name string) string{ // gabisa karna hruf kecil
		return "bye " + name
}
