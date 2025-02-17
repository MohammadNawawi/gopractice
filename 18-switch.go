package main

import "fmt"

func main() {
	name := "Naw"

	// CARA NORMAL
	// switch name {
	// case "Naw":
	// 	fmt.Println("Hello Naw")
	// 	fmt.Println("How are you today?")
	// case "Moh":
	// 	fmt.Println("Hello Moh")
	// 	fmt.Println("Are you okay?")
	// default:
	// 	fmt.Println("Who are you?")
	// }

	//SHORT STATEMENT
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Terlalu panjang!")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	//SWITCH TANPA CONDISI
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Terlalu panjang!")
	case length < 3:
		fmt.Println("Terlalu pendek!")
	default:
		fmt.Println("Nama sudah benar")
	}
}