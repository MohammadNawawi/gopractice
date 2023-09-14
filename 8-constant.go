package main

import "fmt"

func main() {
	const firstName string = "Mohammad"
	const lastName = "Nawawi"
	const age = 99

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	// firstName = "Monarc" Tidak bisa mengassign nalai baru ke const

	// Atau bisa menggunakan multiple const seperti pada dibawah ini
	// const (
	// 	firstName string = "Mohammad"
	// 	lastName = "Nawawi"
	// 	age = 99
	// )
}