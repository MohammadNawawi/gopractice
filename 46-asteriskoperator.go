package main

import "fmt"

type Addres struct {
	city     string
	province string
	national string
}

func main() {
	var address1 Addres = Addres{"Bangkalan", "Jawa Timur", "Indonesia"}
	var address2 *Addres = &address1
	address2.city = "New York"
	fmt.Println(address1) // Tidak Berubah
	fmt.Println(address2) // Berubah

	address2 = &Addres{"Bangka-lan", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}
