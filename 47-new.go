package main

import "fmt"

type NewAddress struct {
	city     string
	province string
	national string
}

func main() {
	var alamat1 *NewAddress = new(NewAddress)
	var alamat2 *NewAddress = alamat1

	alamat2.national = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
