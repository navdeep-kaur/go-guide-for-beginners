package main

import "fmt"

// START OMIT
func main() {
	v1 := []int{1, 2, 3}
	v1 = append(v1, 4)

	fmt.Println(len(v1)) // 4
	fmt.Println(v1[1])   // 2
	fmt.Println(v1[1:3]) // [2 3]

	v2 := make(map[string]string)
	v2["hello"] = "world"

	v3, ok := v2["hello"]
	fmt.Println(ok, v3) // true world
} // HL12
//END OMIT
