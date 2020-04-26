package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

// Go passes args to a function as copies
// Therefore passing a normally instantiated "Saiyan" struct to the "Super" function
// will not update its power value, and will print 9000
// This code demonstrates how to achieve the behaviour we require: 

func main() {
	// & = the "address of" operator
	// here, the value "goku" is now an pointer/address to the value
	goku := &Saiyan{ "Goku", 9000 }

	// Pass the address as the value
	Super(goku)

	// => 19000
	fmt.Println(goku.Power)
}

// Go passes variables to functions as copies.
// The "*" dictates that it expects an pointer to value of type "Saiyan"
func super(s *Saiyan) {

	// This now updates the original
	s.Power += 10000;
}