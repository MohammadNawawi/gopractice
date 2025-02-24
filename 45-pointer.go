package main

import "fmt"

type Address struct {
	city     string
	province string
	national string
}

func main() {
	////Pass by value
	//address1 := Address{"Bangkalan", "Jawa Timur", "Indonesia"}
	//address2 := address1
	//
	//address2.city = "New York"
	//
	//fmt.Println(address1) // Tidak Berubah
	//fmt.Println(address2) // Berubah

	//Pointer
	address1 := Address{"Bangkalan", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.city = "New York"

	fmt.Println(address1) // IkutBerubah
	fmt.Println(address2) // Berubah
}
