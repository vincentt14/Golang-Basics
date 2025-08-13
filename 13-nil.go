package main

import "fmt"

type Map map[string]string

func NewMap(name string) Map {
	if name == "" {
		return nil
	} else {
		return Map{
			"name": name,
		}
	}
}

func main13() {
	data := NewMap("")

	if data == nil{
		fmt.Println("data map masih kosong")
	}else {
		fmt.Println(data["name"])
	}	
}