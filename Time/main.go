package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// Standard time used for formatting as in the documentation
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	// Creating a date
	createdDate := time.Date(2021, time.December, 21, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	// Standard time used for formatting as in the documentation
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))

}
