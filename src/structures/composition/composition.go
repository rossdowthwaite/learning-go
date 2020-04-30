// Composition: The act of including one structure into another (trait/mixin)

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

// Giving the struct the field type of "*Person" with a field name
// allows us to access the fields & functions of "Person" 
type Saiyan struct {
	*Person
	Power int
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku", 10000 },
		Power: 9001,
	}
	goku.Introduce()

	// Can directly access the fields of the composed type
	fmt.Println(goku.Name) // => "Goku"
	fmt.Println(goku.Age) // => 10000
	// same as:
	fmt.Println(goku.Person.Name) // "Goku"
	fmt.Println(goku.Person.Age) // => 10000
}

// Assigning function to Person struct
// Person is the reciever of this function
func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}



