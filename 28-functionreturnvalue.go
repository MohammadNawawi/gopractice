package main

import "fmt"

func getHello(firstName string, lastName string) string {
	return "Hello World " + firstName + " " + lastName
}
func main() {

	result := getHello("Mohammad", "Nawawi")
	fmt.Println(result)
}