package main

import "fmt"

func main() {
	var ersterString string = "Hallo"
	var zweiterString string = "Welt"
	fmt.Println(ersterString)
	fmt.Println(zweiterString)
	fmt.Println()

	var dritterString *string = &ersterString
	var vierterString *string = &zweiterString
	fmt.Println("1:", &ersterString)
	fmt.Println("2:", &zweiterString)
	fmt.Println("3:", dritterString)
	fmt.Println("$:", vierterString)
	fmt.Println()

}
