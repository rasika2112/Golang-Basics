package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our dance classes!")
	fmt.Println("Please rate our dance classes!")

	// Initialising the reader
	reader := bufio.NewReader(os.Stdin)

	// Taking the user input until next line
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	// This will show error to convert 4 from string to float with new line, i.e '4\n' cannot be converted to float
	// numRating, err := strconv.ParseFloat(input, 64)

	// This will convert 4 string to float as we are trimming the extra spaces
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
