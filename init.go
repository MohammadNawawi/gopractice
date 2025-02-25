package main

import (
	"fmt"
	"gopractice/database"
	_ "gopractice/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
