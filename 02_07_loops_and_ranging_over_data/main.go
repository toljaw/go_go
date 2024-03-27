package main

import (
	"fmt"
	"log"
)

func main() {
	for_loop()
	fmt.Println()
	for_slice()
	for_map()
	for_string()
	for_struct()
}

func for_loop() {
	for i := 0; i <= 10; i++ { // 0 and 10 included
		log.Println(i)
	}
}

func for_slice() {
	animals := []string{
		"dog",
		"fish",
		"horse",
		"cow",
		"cat",
	}

	for _, animal := range animals { // if variable starts with "_", the compiler will ignore that it is unused
		log.Println(animal)
	}
}

func for_map() {
	animals := make(map[string]string)

	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"
	animals["horse"] = "Charly"
	animals["bird"] = "Tshirpy"

	for animalType, animal := range animals {
		fmt.Println(animalType, animal)

	}
}

func for_string() {
	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine { // returns the byte value of each char inside the string
		log.Println(i, ":", l)
	}
}

func for_struct() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@browm.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

	for _, l := range users {
		fmt.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
