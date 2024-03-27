package main

import "log"

func main() {
	ifElse()
	ifElseIf()
	switchStatement()
}

func ifElse() {
	isTrue := true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "dog" {
		log.Println("Cat is cat.")
	} else {
		log.Println("cat is not cat.")
	}
}

func ifElseIf() {
	myNum := 100
	isTrue := true

	if myNum > 99 && isTrue == false { // "isTrue == false" = "!isTrue"
		log.Println("myNum is greater then 99 and isTrue is false.")
	} else if myNum == 98 || isTrue == true {
		log.Println("You don't know!")
	}
}

func switchStatement() {
	myVar := "bird"
	switch myVar { // as soon one matches the case, the check will stop running
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}
}
