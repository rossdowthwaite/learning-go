package main

// Imports are strict - if an import is declared but not used, it will not compile.
import (
    "fmt"
    "os"
)

func main() {
	
	// Checks the len is two, becuase first item is always the path the executable.
	// Therefor if lenth is 1, no arg has been provided.
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	fmt.Println("Hello,", os.Args[1]);
	fmt.Println("Path:", os.Args[0]);
}

// Usage: go run args.go world