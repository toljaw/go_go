package main

import (
	"log"
	"sort"
)

func main() {
	//var myString string
	//var myInt int
	//
	//myString = "Hi"
	//myInt = 11
	//
	//mySecondString := "another string"
	//
	//log.Println(myString, myInt, mySecondString)

	myMap := make(map[string]string)

	myMap["dog"] = "Samson" // can store as many things as you want inside the map
	myMap["other dog"] = "Cassie"
	myMap["dog"] = "Fido" // you can overwrite content

	log.Println(myMap["dog"])
	log.Println(myMap["other dog"])

	second()

	third()

	slice()
}

func second() {
	myMap := make(map[string]int)

	myMap["First"] = 12
	myMap["Second"] = 24

	log.Println("First integer:", myMap["First"])
	log.Println("Second integer:", myMap["Second"])
}

type User struct {
	FirstName string
	LastName  string
}

func third() {
	myMap := make(map[string]User) // zuvor erstellte struct wird hier in map eingesetzt

	me := User{
		FirstName: "Anatoli",
		LastName:  "Weikum",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)
}

func slice() {
	var mySlice []string

	mySlice = append(mySlice, "Sandy")
	mySlice = append(mySlice, "Louise", "Anatoli")

	var myIntSlice []int

	myIntSlice = append(myIntSlice, 52, 45, 3, 5, 7, 13)
	sort.Ints(myIntSlice)

	numbers := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(mySlice)
	log.Println(myIntSlice)
	log.Println(numbers)
	log.Println("The first 4 numbers in 'numbers' are:", numbers[0:4]) // last index in slice not included
	log.Println(names)
}
