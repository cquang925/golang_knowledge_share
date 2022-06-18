package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	// creating an empty map using shorthand syntax
	myMap := make(map[string]string)

	// appending map myMap
	myMap["dog"] = "Samson"
	myMap["cat"] = "Sophie"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	// example of overwriting key
	myMap["dog"] = "Milo"
	log.Println(myMap["dog"])

	// -------------
	// example of creating map with different types for key/value
	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])

	// -------------
	// maps can store anything such as creating a type and storing type in the map
	// example of creating map and assigning to a struct User
	myMap3 := make(map[string]User)

	// creates a User named person
	person := User {
		FirstName: "Charles",
		LastName: "Quang",
	}

	// initializes myMap3 and assigns to User person
	myMap3["me"] = person

	// prints myMap3 with key of "me"
	log.Println(myMap3["me"])
	// prints field of key "me" in map myMap3
	log.Println(myMap3["me"].LastName)
}