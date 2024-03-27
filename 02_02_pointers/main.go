package main

import "log"

func main() {
	var myString string = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) // & means reference to a variable
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue // * adds a pointer type to a variable
}
