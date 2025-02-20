package main

import "fmt"

func logging() {
	fmt.Println("Function Logging")
}

func runApp() {
	defer logging()
	fmt.Println("Function RunApp")
}
func main() {
	runApp()
}
