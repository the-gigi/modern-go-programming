package vet_demo

import "fmt"

func foo() {
	return;

	// Unreachable code
	fmt.Printf("This will never be printed")

}
