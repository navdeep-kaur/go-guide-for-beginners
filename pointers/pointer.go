package main

import (
	"fmt"
)

func getShape(input int) *string {
	var result *string

	if input < 0 {
		return result
	}
	if input == 0 {
		s := "Circle"
		return &s
	}

	if input == 1 {
		s := "Square"
		return &s
	}
	return result
}
func main() {
	//shape := getShape(0)
	var shapeAddress *string
	shapeAddress = getShape(1)

	fmt.Println(shapeAddress)
	fmt.Println(*shapeAddress)

	//fmt.Println("")
	//myString := "a string value"
	//fmt.Println(myString)
	//
	//myInteger := 1
	//fmt.Println(myInteger)

	//ints := make([]int, 4)

	defer cleanResources()

	v1 := []int{1, 2, 3}
	v1 = append(v1, 4) // 1, 2, 3, 4

	fmt.Println(len(v1)) // 4
	fmt.Println(v1[1])   // 2

	newSlice := v1[1:3] // [2 3]

	newSlice[0] = 5
	fmt.Println(v1) // [1, 5, 3, 4]

	v2 := make(map[string]string)
	v2["hello"] = "world"

	v3, ok := v2["hello"]
	fmt.Println(ok, v3) // true world

	var dog Animal = &Dog{
		Name: "Cooper",
	}
	var cat Cat = "myCat"
	var catInterface Animal = cat
	cat.Walk()

	fmt.Println(dog.Speak())
	fmt.Println(catInterface.Speak())

	var Joan *string

	if Joan != nil {
		fmt.Println(&Joan) // --> ""
	}

	var abc ExtendedAnimal
	abc.Speak()

}

func cleanResources() {

}

type Animal interface {
	Speak() string
	Walk() error
	ChangeName(newName string)
}

type ExtendedAnimal interface {
	Animal
	Dance() string
}

type WalkingStruct struct {
}

type Dog struct {
	WalkingStruct
	Name string
}

type Cat struct {
	WalkingStruct
	Name string
}
type CatInt int8

func (w WalkingStruct) Walk() error {
	return nil
}

// here we define Speak as a method on Dog
// it takes a pointer receiver, but you
// could also remove the *
func (d Dog) Speak() string {
	return fmt.Sprintf("I am %s", d.Name)
}

func (d Dog) ChangeName(newName interface{}) {

	s := newName.(string)
	d.Name = s
}

//func (d *Dog) ChangeName(newName string) {
//	d.Name = newName
//}

func (c Cat) Speak() string {
	return string(c)
}

func (c Cat) ChangeName(newName string) {
	c = newName
}
