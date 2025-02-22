package main

import "fmt"

func random() any {
	return 1
}
func main() {
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	//
	//var resultInt int = result.(int) //Akan panic
	//fmt.Println(resultInt)

	//Penggunaan type assertions yang benar apabila ingin merubah type
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}

}
