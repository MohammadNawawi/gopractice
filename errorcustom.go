package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id is empty"}
	}

	if id != "Naw" {
		return &notFoundError{Message: "id is not Naw"}
	}

	//ok
	return nil
}
func main() {
	err := SaveData("Naw", nil)
	if err != nil {
		//If
		//terjadi error
		//if validationErr, ok := err.(*validationError); ok {
		//	fmt.Println("Validation Error", validationErr.Error())
		//} else if notFoundErr, ok := err.(*notFoundError); ok {
		//	fmt.Println("Not found Error", notFoundErr.Error())
		//} else {
		//	fmt.Println("Unknown Error", err.Error())
		//}

		//Switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found Error", finalError.Error())
		default:
			fmt.Println("Unknown Error", finalError.Error())
		}
	} else {
		//Sukses
		fmt.Println("Successfully saved data")
	}
}
