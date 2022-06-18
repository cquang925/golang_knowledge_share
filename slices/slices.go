package main

import (
	"log"
	"sort"
)

func main() {
	// example creating a slice
	var mySlice []string

	// example appending slice
	mySlice = append(mySlice, "Charles")
	mySlice = append(mySlice, "Adam")
	mySlice = append(mySlice, "Steve")
	mySlice = append(mySlice, "Mary")

	// prints out slice in order mySlice was appended
	log.Println(mySlice)

	// example on how sort a slice
	sort.Strings(mySlice)
	log.Println(mySlice)

	// -------------------------------
	// exmaple of declaring slice using shorthand
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// example of creating a range of slice from numbers
	log.Println(numbers[0:3])
	log.Println(numbers[5:9])

	rangeSlice := numbers[2:6]
	log.Println(rangeSlice)

}
