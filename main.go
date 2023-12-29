package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "SKN"

	const authorId string = "skn437"

	fmt.Printf("Hello! Your ID is %s\n", authorId)

	fmt.Println("Hello", name, "!")

	var username string

	fmt.Scan(&username)

	fmt.Printf("You username is %s!\n", username)

	fmt.Printf("authorId is %T\n", authorId)

	var myArray [2]string

	myArray[0] = "SKN"
	myArray[1] = "Wang So"

	fmt.Println(myArray)

	var array []string

	array = append(array, "Wang")
	array = append(array, "So")

	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Array Length: %v\n", len(array))

	for i := 0; i <= 7; i++ {
		fmt.Println(i)
	}

	for index, arr := range array {
		fmt.Printf("%v %v\n", index, arr)
	}

	var myNames []string

	myNames = append(myNames, "SKN Shukhan")
	myNames = append(myNames, "Wang So")

	for _, name := range myNames {
		var firstName []string = strings.Fields(name) //* string.Field() returns a "Slice"

		fmt.Printf("First name: %s\n", firstName)
	}
}
