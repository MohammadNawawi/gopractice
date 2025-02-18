package main

import "fmt"

func getFullName() (string, string) {
	return "Mohammad", "Nawawi"
}
func main() {

	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName() //Tidak menggunakan last name
	fmt.Println(firstName)
}