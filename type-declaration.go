package main

import "fmt"

func main6() {
	type NoKTP string

	var myKTP NoKTP = "111111111111"
	fmt.Println(myKTP)
	fmt.Println(NoKTP("222222222222"))
}