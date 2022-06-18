package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

// example of func having pointer to struct myStruct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {

	// example of initializing struct and assigning values
	user := User{
		FirstName: "Charles",
		LastName:  "Quang",
	}

	// printing out created user with undefined field. Age is defaulted 
	log.Println(user.FirstName, user.LastName, user.Age)

	// example of initializing struct and declaring by long hand and shorthand
	// longhand ex
	var myVar myStruct
	myVar.FirstName = "Dale"

	// shorthand ex
	myVar2 := myStruct {
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	// example using function to call variable.
	// func printFristName has pointer to struct myStruct
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

}
