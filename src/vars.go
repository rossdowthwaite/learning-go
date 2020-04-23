package main

import (
	"fmt"
)

func main() {
	// Declare and assign
	var powerOne int = 9000

	// Declare and assign and infer the type
	powerTwo := 8000

	// Variables cannot be defined twice in the same scope. 
	// Doing so will cause a compilation error.
	// So to re-assign a variable use "=":
	powerTwo = 10000

	fmt.Printf("It's over %d\n", powerOne)
	fmt.Printf("It's over %d\n", powerTwo)

	defineMoreVariables()
	potentialCompilerCrapOut()
}

func defineMoreVariables() {
	thing := 1000
	fmt.Printf("this is a %d\n", thing)

	// One would expect the compiler to crap out here due to re-declaring "thing" in the same scope
	// However this is not the case, as the new var "anotherThing" is a new variable
	// BUT! the type "thing" cannot be changed 
	anotherThing, thing := "beaver", 100000
	fmt.Printf("a %s has over %d hairs per square cm\n", anotherThing, thing);
}

func potentialCompilerCrapOut() {
	cheese, buscuit := "Cheddar", "Jacobs cracker"

	// NO! compiler will crap out because "buscuit" is not used!
	fmt.Printf("%d is a type of cheese\n", cheese)

	// must be used
	fmt.Printf("%d is a type of cracker\n", buscuit)
}