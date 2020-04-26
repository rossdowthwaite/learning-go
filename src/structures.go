package main

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
}