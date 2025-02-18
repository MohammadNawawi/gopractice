package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	
	firstName = "Mohammad"
	middleName = "Nawawi"
	lastName = "Nawawi"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}