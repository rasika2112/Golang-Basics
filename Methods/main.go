package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")

	user := User{"Rasika", "rasika@gmail.com", true, 21}
	fmt.Println(user)
	user.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Method to get status of user
func (u User) GetStatus() {
	fmt.Println("Is User Active?: ", u.Status)
}
