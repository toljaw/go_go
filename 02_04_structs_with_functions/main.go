package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string { // like a method in oop
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Johny"

	myVar2 := myStruct{
		FirstName: "Donny",
	}

	log.Println("myVar is set to", myVar.printFirstName)
	log.Println("myVar2 is set to", myVar2.printFirstName)
}
