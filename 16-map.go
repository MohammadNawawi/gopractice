package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Mohammad",
		"address": "Martajasah",
	}

	person["hobby"] = "Sport"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["hobby"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar GoLang"
	book["author"] = "Naw"
	book["ups"] = "Wrong"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}