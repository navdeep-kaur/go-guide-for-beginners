package main

import "fmt"

// START OMIT

var someRandomVariable = "Navdeep is Awesome"

func myFunc() string {
	return "Hello World"
}
func main() {
	v1 := "some variable"

	var v2 = "some other variable"
	v1 = "now v1 has a different value"
	v2 = " same with v2"
	fmt.Print(v1, v2)

	var myFuncResponse = myFunc()
	fmt.Println(myFuncResponse)

	fmt.Println(someRandomVariable)
} // HL12
//END OMIT
