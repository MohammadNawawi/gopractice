package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4} //Apabila kita memiliki slice/array dan ingin mengirimkannya maka kita memberikan "..."
	fmt.Println(sumAll(numbers...))
}
