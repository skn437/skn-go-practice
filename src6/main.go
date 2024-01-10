package main

import "fmt"

// * More Slice Practices
func main() {
	var names []string = []string{
		"SKN",
	}

	names = append(names, "Wang", "So", "Shukhan") //* append can take multiple values to inject in slices

	fmt.Printf("Names: %v \n", names)

	//* Slicing The slices
	var newNames []string = names[1:3]

	fmt.Printf("Names: %v \n", names)
	fmt.Printf("New Names: %v \n", newNames)

	newNames = newNames[:1]

	fmt.Printf("Names: %v \n", names)
	fmt.Printf("New Names: %v \n", newNames)
}
