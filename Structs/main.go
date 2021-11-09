package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

// +%v in printf will print the keys also
func main() {
	// Struct is a value type
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)

	// Cannot iterate over a struct
	fmt.Printf("%+v", alex)
}
