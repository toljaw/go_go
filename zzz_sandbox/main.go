package main

import "fmt"

//import (
//	"fmt"
//)
//
//func main() {
//	var quarters int
//	var dimes int
//	var nickles int
//	var pennies int
//
//	fmt.Println("Please insert quarters")
//	_, err := fmt.Scanln(&quarters)
//	if err != nil {
//		fmt.Println("Error during conversion", err)
//		return
//	}
//	fmt.Println("Please insert dimes")
//	_, err = fmt.Scanln(&dimes)
//	if err != nil {
//		fmt.Println("Error during conversion", err)
//		return
//	}
//	fmt.Println("Please insert nickles")
//	_, err = fmt.Scanln(&nickles)
//	if err != nil {
//		fmt.Println("Error during conversion", err)
//		return
//	}
//	fmt.Println("Please insert pennies")
//	_, err = fmt.Scanln(&pennies)
//	if err != nil {
//		fmt.Println("Error during conversion", err)
//		return
//	}
//
//	fmt.Println(quarters, dimes, nickles, pennies)
//	fmt.Printf("Der Typ von quarters ist %T\n", quarters)
//	fmt.Printf("Der Typ von dimes ist %T\n", dimes)
//	fmt.Printf("Der Typ von nickles ist %T\n", nickles)
//	fmt.Printf("Der Typ von pennies ist %T\n", pennies)
//}

func main() {
	var quarters int
	var dimes int
	var nickles int
	var pennies int

	fmt.Println("Please insert quarters")
	fmt.Scanln(&quarters)
	fmt.Println("Please insert dimes")
	fmt.Scanln(&dimes)
	fmt.Println("Please insert nickles")
	fmt.Scanln(&nickles)
	fmt.Println("Please insert pennies")
	fmt.Scanln(&pennies)

	fmt.Println(quarters, dimes, nickles, pennies)
	fmt.Printf("Der Typ von quarters ist %T\n", quarters)
	fmt.Printf("Der Typ von dimes ist %T\n", dimes)
	fmt.Printf("Der Typ von nickles ist %T\n", nickles)
	fmt.Printf("Der Typ von pennies ist %T\n", pennies)
}
