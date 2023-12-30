package main

import (
	"fmt"
	"skn-go-practice/src/pkgs/functions"
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

	var count uint8 = 7

	if count < 10 {
		fmt.Printf("Count: %v\n", count)
	} else {
		fmt.Println("Hello!")
	}

	var newCount uint8 = 3

	for newCount > 0 { //* for loop can do more than "for" in Go. It can also be used as while loop used in other languages
		fmt.Printf("newCount: %v\n", newCount)
		newCount--
	}

	var email string

	fmt.Print("Enter an email: ")
	fmt.Scan(&email)

	if strings.Contains(email, "@") && len(email) > 1 {
		fmt.Printf("You've entered %s\n", email)
	} else {
		fmt.Println("Invalid Email!")
	}

	var books uint8 = 3

	switch books { //* Go switch doesn't need "break" just like in TypeScript
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("Nothing!")
	}

	fmt.Printf("Get New Array Result: %v\n", functions.GetNewArray(14))

	var num1, num2, num3 uint8 = functions.GetMultipleReturns("SKN")

	fmt.Printf("%v %v %v\n", num1, num2, num3)

}
