package main

import(
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku", 10000},
		Power: 1000,
	}

	goku.Introduce() // => "Hi, I'm overloading person and my name is Goku. Ya!"
	goku.Person.Introduce() // => "Hi, I am a person called Goku"
}

func (p *Person) Introduce() {
	fmt.Println("Hi, I am a person called %s", p.Name)
}

// We can overwrite functions of a composed type
func (s *Saiyan) Introduce() {
	fmt.Println("Hi, I'm overloading person and my name is %s. Ya!\n", s.Name)
}
