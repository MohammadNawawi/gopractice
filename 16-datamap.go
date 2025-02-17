package main

import "fmt"

func main() {
	person := map[string]string{"name": "Nawawi", "title": "Programmer"}
	fmt.Println(person["name"])
	fmt.Println(person)
}