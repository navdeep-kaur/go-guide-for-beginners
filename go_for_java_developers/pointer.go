package main

import "fmt"

// START OMIT
func main() {
	v1 := "some variable"
	// v2 is now a pointer to v1
	var v2 *string = &v1
	fmt.Print(v2) // prints the memory address of v1
} // HL12
//END OMIT
