package main

import "fmt"

func main() {
	var name string

	name = "Mohammad"
	fmt.Println(name)

	name = "Nawawi"
	fmt.Println(name)

	var frindName = "Jono"
	fmt.Println(frindName)

	var age = 99
	fmt.Println(age)

	country := "Mexico" // ':' merupakan var dimana ini diguanakan pada awal deklarasi saja, selanjutnya langsung menggunakan = saja
	fmt.Println(country)

	var (
		firstName = "Mohammad"
		lastName = "Nawawi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}