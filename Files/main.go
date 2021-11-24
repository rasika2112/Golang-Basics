package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")

	content := "This content will be written in the file"

	// Creating a text file in the same directory
	file, err := os.Create("./file.txt")

	// Catching error and stopping thefurther execution
	checkNilErr(err)

	// Writing content to the file
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is: ", length)

	// Reading the file content
	readFile("./file.txt")
	defer file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside file is\n", string(databyte))
}

// Function to catch error as it is used many times
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
