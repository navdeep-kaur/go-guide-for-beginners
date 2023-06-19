package main

import (
	"errors"
	"fmt"
)

// START OMIT
func test(input int) error {
	if input < 0 {
		return errors.New("less than zero")
	}
	return nil
}
func main() {
	err := test(-1)
	if err != nil {
		fmt.Print(err)
	}
} // HL12
//END OMIT
