package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["white"] = "FFFFFF"

	// delete(colors, "white")
	colors := map[string]string{
		"red":   "FF0000",
		"blue":  "00FF00",
		"green": "0000FF",
		"white": "FFFFFF",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, ":", hex)
	}
}
