package main

import "fmt"

// START OMIT
type Animal interface {
	Speak() string
}
type Dog struct {
	Name string
}

// here we define Speak as a method on Dog
// it takes a pointer receiver, but you
// could also remove the *
func (d *Dog) Speak() string {
	return fmt.Sprintf("I am %s", d.Name)
}
func main() {
	var dog Animal = &Dog{
		Name: "Cooper",
	}

	fmt.Println(dog.Speak())
} // HL12
//END OMIT
