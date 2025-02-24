package main

import "fmt"

type PointerFunctionAddress struct {
	city     string
	province string
	national string
}

func ChangeCountryToId(address *PointerFunctionAddress) {
	address.city = "Bangkal"
}

func main() {
	var address PointerFunctionAddress = PointerFunctionAddress{}
	ChangeCountryToId(&address)
	fmt.Println(address)
}
