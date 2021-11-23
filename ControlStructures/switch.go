package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SwitchCase() {
	fmt.Println("Switch Case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // rand will generate no. between 0 and 5
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
		// fallthrough will also run the next case
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("Dice value is not 1 or 6")

	}
}
