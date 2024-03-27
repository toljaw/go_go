package main

import (
	"log"
	"time"
)

//var s = "seven" // variable is declared outside any function - it has global scope
//
//func main() {
//	var s2 = "six" // variable is declared inside a function - it has local scope and can only be used inse the
//	// function
//
//	log.Println("s is ", s)
//	log.Println("s2 is ", s2)
//
//	saySomthing("blubb")
//}
//
//func saySomthing(s3 string) (string, string) { // if you use "s" here, then the local s will be used
//	log.Println("s from the saySomething func is", s) // because in line 17 is not s declared,
//	// global s will be used
//	return s, "World"
//}

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user1 := User{
		FirstName: "Anatoli",
		LastName:  "Weikum",
		// you don't have to use all of them
	}
	log.Println("Hi, my name is", user1.FirstName, user1.LastName)
}

// if you declare a variable like "example", it is available only in this file
// if you declare a variable like "Example" with a capital letter,
//it is available also in other files - same for functions - a little bit like classes
