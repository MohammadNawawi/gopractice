package main

import "fmt"

type Cust struct {
	name    string
	address string
	age     int
}

func (cust Cust) sayHello(name string) {
	fmt.Println("Hello ", name, "my name is : ", cust.name)
}

func main() {
	fmt.Println("Function main")

	var customer Cust
	customer.name = "Jack"
	customer.address = "Jalan Buntu"
	customer.age = 25
	fmt.Println(customer)
	fmt.Println(customer.name)
	fmt.Println(customer.address)
	fmt.Println(customer.age)

	customer.sayHello("Naw")
}
