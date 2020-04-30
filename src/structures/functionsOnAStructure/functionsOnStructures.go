// Associating a method with a structure:

package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

func main() {
	goku := &Saiyan{ "Goku", 9000 }
	
	// No args needed here
	goku.Super() 
	
	fmt.Println(goku.Power) // => 10000  
}

// Type *Saiyan is the "reciever" of the Super method.
func (s *Saiyan) Super() {
	s.Power += 10000
}



