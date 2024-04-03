package main

import (
	"go_go/02_09_channels/helpers"
	"log"
)

const numPool = 10 // pool of numbers from which random numbers will be created from

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
