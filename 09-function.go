package main

import "fmt"

func main9(){
	sayHello()
	sayHelloTo("Eko", "Vincent")

	result := getHello("Vincent")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// multiple return value wajib ditangkap sama variabel. kalau ga mau tangkap pakai _
	_, namaMu := getFullName()
	fmt.Println(namaMu)

	a, b, c := getCompleteNames()
	fmt.Println(a, b, c)

	total1 := sumAllVarArg(10, 10, 10, 10, 10)
	total2 := sumAllSlice([]int{10, 10, 10, 30})
	fmt.Println(total1, total2)

	// kalau udah terlanjur punya slice, kirim ke paramnya harus kasi ... agar auto konversi ke varargs dari slice
	numbers := []int{10, 10, 10}
	fmt.Println(sumAllVarArg(numbers...))

	// function as a value
	goodbye := getGoodbye
	fmt.Println(goodbye("Cent"))

	// function as parameter
	sayHelloWithFilter("anjing", spamFilter)

	// type di bawah

	// anonim function
	registerUser("Vincent", func(name string) bool {
		return name == "anjing"
	})
	
	// ATAU param ke 2 masukin ke variabel blaclist := paste ke
	// registerUser("Vincent", blacklist)

	// recursive function - CASE: factorial 10
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}

func sayHello(){
	fmt.Println("Hello")
}

func sayHelloTo(nama1 string, nama2 string){
	fmt.Println("Hello bang", nama1, nama2)
}

func getHello (name string)string {
	return "Hellow " + name
}

// multiple returning func
func getFullName ()(string, string){
	return "Eko", "Khanedy"
}

// named return value
func getCompleteNames() (name1, name2, name3 string){
	name1 = "baba"
	name2 = "bubu"
	name3 = "bobo"

	return name1, name2, name3
}

// variadic function / Variable Argument
func sumAllVarArg(numbers ...int)int {
	total := 0
	for _, number := range numbers{
		total += number
	}

	return total
}

func sumAllSlice(numbers []int)int {
	total := 0
	for _, number := range numbers{
		total += number
	}

	return total
}

// function variabel = function as a value
func getGoodbye (name string)string {
	return "Goodbye " + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string)string){
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing"{
		return "..."
	}else {
		return name
	}
}

// function type declaration -> KAYAK TYPESCRIPT
type Filter func(string)string

func sayHelloWithFilterType(name string, filter Filter){
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

//anonim function
type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You are Blocked", name)
	}else {
		fmt.Println("Welcome", name)
	}
}

// factorial loop
func factorialLoop(value int)int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int)int {
	if (value == 1){
		return 1
	}else {
		return value * factorialRecursive(value -1)
	}
}