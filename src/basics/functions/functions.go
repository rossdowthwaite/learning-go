package main

import (
	"fmt"
)

func main() {
	value, exists := power("goku");

	// Sometimes only care about one variable: use "_"
	// The "_" remains unassigned and can be reused multiple times
	_, thing := son("gohan")

	fmt.Printf("%s %t %f", value, exists, thing)
}

// Defines a function called add, that returns an "int"
// shorthand params when they share the same type
func add(a, b int) int {
	return a + b
}

// Defines a function called add, that returns an "int" and "bool"
func power(name string) (int, bool) {
	return 10000, true
}

func son(name string) (string, bool) {
	return "annoying child", true
}