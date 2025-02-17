package main

import "fmt"

func main() {
	name := "Naw"

	if name == "Naw" {
		fmt.Println("Hello Naw")
	}else if name == "Moh" {
		fmt.Println("Hello Moh")
	}else{
		fmt.Sprintln("Who are you?")
	}

	// CARA NORMAL
	// length := len(name)
	// if length > 5  {
	// 	fmt.Println("Terlalu panjang!")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	// SHORT STATEMENT
	if length := len(name) ; length > 5 {
		fmt.Println("Terlalu panjang!")
	}else{
		fmt.Println("Nama sudah benar")
	}
}