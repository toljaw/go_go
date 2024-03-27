package main

import "fmt"

func main() {
	var name string
	fmt.Println("Wie heißt du?")
	fmt.Scanln(&name)

	fmt.Println("Hallo", name, "!")
	fmt.Println(name)
	fmt.Println(&name)
}
