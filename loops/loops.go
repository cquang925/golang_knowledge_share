package main

import "log"

func main() {

	// for loop example, sets variable i to 0, checks to see if i is greater or equal to 5, increments i by 1
	// once i is greater than 5 loop will end
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	// ---------------
	// slice of strings
	animals := []string{"dog", "fish", "cat", "horse"}

	// example of for loop with range
	// below will print index of animal and the value in the slice
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// example how to range over slice without printing index
	for _, animal := range animals {
		log.Println(animal)
	}

	// below example will only print the index and not value
	for animal := range animals {
		log.Println(animal)
	}

	// ---------------
	// looping over map range example
	pets := make(map[string]string)
	pets["dog"] = "Barky"
	pets["cat"] = "Sophie"
	pets["horse"] = "Ed"

	// example will print the values of the map
	for _, pet := range pets {
		log.Println(pet)
	}

	// example will print key and value of map
	for petType, pet := range pets {
		log.Println(pet, "is a", petType)
	}

	// ---------------
	// looping over custom type
	type User struct {
		FirstName string
		LastName string
		Age int
	}

	var users []User
	users = append(users, User{"Charles", "Quang", 39})
	users = append(users, User{"John", "Smith", 22})
	users = append(users, User{"Allison", "Davis", 30})

	// example loops over custom type User and ignores index
	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Age)
	}
}