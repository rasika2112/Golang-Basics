package main

import "fmt"

func IfElse() {
	fmt.Println("If Else in golang")

	count := 21
	var result string

	if count < 10 { // { bracket should not start from next line
		result = "Regular user"
	} else if count > 10 {
		result = "Watch out"
	} else {
		result = "Exact count"
	}

	fmt.Println(result)
}
