package main

import (
	"fmt"
)

// Example of a simple structure
type Saiyan struct {
	Name string
	Power int
}

func main() {	
	// Simple assignment 
	saiyan1 := Saiyan{
		Name: "Goku",
		Power: 9000,
	}

	//  Fields do not need to be set when assigned
	saiyan2 := Saiyan{}

	// or...
	saiyan3 := Saiyan{ Name: "Goku" }
	saiyan3.Power = 12

	// or even shorter
	saiyan4 := Saiyan{"Vegeta", 9000}

	fmt.Printf(saiyan1.Name)
	fmt.Printf(saiyan2.Name)
	fmt.Printf(saiyan3.Name)
	fmt.Printf(saiyan4.Name)
}