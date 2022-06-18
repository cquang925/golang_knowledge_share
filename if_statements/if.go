package main

import "log"

func main() {

	// comment out line 8 and uncomment line 9 to check on difference in below example
	// isTrue := true
	isTrue := false

	// example using if statement to check if variable is true and if so to print line
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// example checking if variable equals to something
	cat := "cat"

	if cat == "cat" {
		log.Println("cat is a cat")
	} else {
		log.Println("cat is NOT a cat")
	}

	// else if statements to check if variable is set to something specfic. *better to use switch
	pet := "cat"

	if pet == "cat" {
		log.Println("your pet is a", pet)
	} else if pet == "dog"{
		log.Println("your pet is a", pet)
	} else if pet == "bird" {
		log.Println("your pet is a", pet)
	} else {
		log.Println("your pet is unknown")
	}

	// change animal and number variables to see what prints on below if statements
	animal := "horse"
	number := 98

	// if statement example where both have to be true to print first line else prints second line
	if number > 99 && animal == "horse" {
		log.Println("number is greater than 99 and animal is horse")
	} else {
		log.Println("number is not greater than 99 and/or animal is not a horse")
	}

	// if statement where one has to be true to print first line, if both are false will print else statement
	if number > 99 || animal == "horse" {
		log.Println("number is greater than 99 or animal is horse")
	} else {
		log.Println("number is not greater than 99 and animal is not a horse")
	}

}