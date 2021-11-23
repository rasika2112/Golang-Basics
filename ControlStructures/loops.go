package main

import "fmt"

func Loops() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Monday", "Tuesday"}

	fmt.Println(days)

	// not used much
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// i is the index
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	val := 1
	for val < 11 {

		if val == 2 {
			goto l
		}
		fmt.Println("Value is: ", val)
		val++
	}

l:
	fmt.Println("Jumped to label using goto")
}
