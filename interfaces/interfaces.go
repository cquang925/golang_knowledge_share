package main

import "fmt"

/*
Example on creating interface type. 
Interface definition will have list of functions that satisfy type Animal.
Below example creates interface named Animal and two separate struct types named Dog and Gorilla.
main() instaniates structs
functions listed in interface definitions will need to be created for the instantiated structs and function needs to point to the structs
function is created that calls on the interface which then calls on the functions of the struct type
*/

// example creating type interfaced named Animal that accepts functions Says and NumberOfLegs
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// creates type structure named Dog with two variables
type Dog struct {
	Name  string
	Breed string
}

// creates type structure named Gorilla with three variables
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	// instantiates type Dog named dog
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	// calls PrintInfo function that receives dog
	PrintInfo(&dog)

	// instantiates type Gorilla named gorilla
	gorilla := Gorilla {"George", "Brown", 28}

	// calls PrintInfo function that receives gorilla
	PrintInfo(&gorilla)

}

// function accepts interface Animal that is stored in variable a
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// function points to type Dog and returns string 'woof'
func (d *Dog) Says() string {
	return "woof"
}

// function points to type Dog and returns int 4
func (d *Dog) NumberOfLegs() int {
	return 4
}

// function points to type Gorilla and returns string 'rawr'
func (g *Gorilla) Says() string {
	return "rawr"
}

// function points to type Gorilla and returns int 2
func (g *Gorilla) NumberOfLegs() int {
	return 2
}
