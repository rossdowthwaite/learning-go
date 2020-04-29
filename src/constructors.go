package main

import (
	"fmt"
)

func main() {

}

func newSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}