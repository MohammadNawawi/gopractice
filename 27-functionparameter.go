package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello World", firstName, lastName)
}
func main() {
	sayHello("Mohammad", "Nawawi")
}