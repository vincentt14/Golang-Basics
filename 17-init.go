package main

import (
	"Basics/database"
	_ "Basics/internal" //dikasi blank identifier (_) supaya initnya jalan aja tapi ga dipakai func yg lainnya
	"fmt"
)

func main17() {
	fmt.Println(database.GetDatabase())
}