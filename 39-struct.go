package main

import "fmt"

type Customer struct {
	name    string
	address string
	age     int
}

func main() {
	var customer Customer
	fmt.Println(customer)

	customer.name = "Jack"
	customer.address = "Jalan Buntu"
	customer.age = 25
	fmt.Println(customer)
	fmt.Println(customer.name)
	fmt.Println(customer.address)
	fmt.Println(customer.age)

	customer2 := Customer{
		name:    "Jack",
		address: "Jalan Buntu",
		age:     25,
	}
	fmt.Println(customer2)

	customer3 := Customer{"Jack", "Jalan Buntu", 25}
	fmt.Println(customer3)

}
