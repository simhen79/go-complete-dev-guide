package main

import "fmt"

type colors map[string]string

func main() {

	// var colors map[string]string

	// kleure := make(map[string]string)

	// kleure["white"] = "#FFFFFF"
	// kleure["rooi"] = "#RED"
	// delete(kleure, "rooi")

	// fmt.Println(colors, kleure)

	colors := colors{
		"red":   "#ff0000",
		"green": "#623467",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c colors) {
	c["red"] = "#RED"
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
