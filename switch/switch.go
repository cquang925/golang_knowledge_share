package main

import "log"

func main() {
	// change value of color to check below switch statement
	color := "red"

	// example of switch checking the variable 'color'. line will print depending on value of color
	switch color {
	case "green":
		log.Println("color is set to green")
	case "blue":
		log.Println("color is set to blue")
	case "red":
		log.Println("color is set to red")
	default:
		log.Println("color is unknown")
	}
}