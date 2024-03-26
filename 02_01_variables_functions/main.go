package main

import "fmt"

// here defined variables are global

func main() { // any go file needs a "main" function - if there is nothing (no datatype) after the parenthesis,
	// there will be no return from the function
	fmt.Println("Hello World.")

	var whatToSay string = "Goodbye, cruel world"
	var i int = 334 // int by default is int64

	fmt.Println(whatToSay)
	fmt.Println()
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething() // the function works even it will be defined later - ":=" will take over
	// the datatype from the variable or function return
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) { // it's possible to return more than one thing - see also line 18 - 20
	return "something", "else"
}
