package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4fhsuj",
	// }

	// All keys should be of same tye and all values too
	// It is reference type
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["green"] = "#fffgff"

	// delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	// Cannot iterate in struct, but we can in maps
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
