package main

import "fmt"

type Man struct {
	Name string
}

// Untuk method direkomendasikan selalu menggunakan pointer *
func (man *Man) Married() {
	man.Name = "Mr" + man.Name
}

func main() {
	isMarried := Man{"John"}
	isMarried.Married()
	fmt.Println(isMarried)
}
