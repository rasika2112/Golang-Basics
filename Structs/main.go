package main

import "fmt"

// no inheritance in golang; No super or parent
// P in capital in Person as it is a class
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// +%v in printf will print the keys also
func main() {
	// Struct is a value type
	alex := Person{"Alex", "Anderson", 21}
	fmt.Println(alex)
	fmt.Printf("Name is %v\n", alex.FirstName)

	// Cannot iterate over a struct
	fmt.Printf("Details are: %+v", alex)
}
