package main

import (
	"fmt"
)

// Struct fields can be of any type, including structures
type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}

func main() {
	gohan := &Saiyan{
		Name: "Gohan",
		Power: 1000,
		Father: &Saiyan {
			Name: "Goku",
			Power: 9001,
			Father: nil, // <= zero value. uninitialized value. NOT undefined
		},
	}

	fmt.Println(gohan.Father.Name)
}
