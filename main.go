package main

import "fmt"

func main() {
	fmt.Println("Maps")
	//var color map[string]string
	//color := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	fmt.Println(colors)
	//fmt.Println(color)
	fmt.Println(colors["green"])
	delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for " + color + " is " + hex)
	}
}
