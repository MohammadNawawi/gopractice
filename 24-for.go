package main

import "fmt"

func main () {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	names := []string{"Nawawi", "Naw", "Mohammad", "Monarc"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, name := range names {
		fmt.Println("Index", index, ":", name)
	}

	// for _, name := range names { //Tidak butuh index diganti jadi _
	// 	fmt.Println(name)
	// }
}