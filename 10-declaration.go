package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var nomorKTP NoKTP = "111111111111"
	var marriedStatus Married = false

	fmt.Println(nomorKTP)
	fmt.Println(marriedStatus)
}