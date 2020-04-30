package main

type Saiyan struct {
	Name string
	Power int
}

// Structures do not have constructors. 
// So instead create a function that returns an instance
// of the desired type
// Note: can, or cannot return a pointer = &Saiyan, return *Saiyan
func newSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name: name,
		Power: power,
	}
}

func main() {
	// Using the "new" function
	// new allocates the memory required by the type
	goku := new(Saiyan)
	goku.name = "Goku"
	goku.power = 9001
	
	// same as the above:
	// Note: this syntax is often prefered as its more readable
	gohan := &Saiyan {
		name: "goku",
		power: 9000,
	}
}