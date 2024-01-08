package main

import (
	"fmt"
	"skn-go-practice/src/pkgs/components"
	"skn-go-practice/src/pkgs/functions"
	"skn-go-practice/src/pkgs/user"
	"skn-go-practice/src/pkgs/utils"
	"sort"
	"strings"
)

func main() {
	var wg = utils.GetWg()

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

	array = append(array, "Wang") //* append doesn't change the slice but instead returns a new array
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

	type User struct {
		firstName string
		lastName  string
		bookCount uint8
	}

	var users = User{
		firstName: "gh",
		lastName:  "hgh",
		bookCount: 4,
	}

	var userArray []User

	userArray = append(userArray, users)

	fmt.Printf("user: %v\n", userArray)

	//for {
	var user2 user.UserType = user.GetUserInput()

	//* Goroutine works fine in loop
	//* If There is no loop then "main" thread will not let another thread to execute code with 'go' keyword
	//* In this case, "Workgroups" have to be used

	wg.Add(5)                     //* `Add` sets the number of Goroutine to wait for
	go user.GetUserConfirmation() //* Goroutine (Concurrence) can be used with 'go' keyword

	fmt.Printf("User: %v\n", user2)

	go user.GetUserCertificate()

	fmt.Println("#############")
	//* Slice Ranges
	var newSlice []string = []string{"Wang", "So", "SKN", "Shukhan"}

	var range1 = newSlice[1:3] //* Slice Range [x:y] means 'From x until y'
	var range2 = newSlice[1:]  //* Slice Range [x:] means 'From x till the end'
	var range3 = newSlice[:3]  //* Slice Range [:y] means 'From the start until y'

	fmt.Printf("New Slice Range1: %v 👍 \n", range1)
	fmt.Printf("New Slice Range2: %v 👍 \n", range2)
	fmt.Printf("New Slice Range3: %v 👍 \n", range3)

	//* `Strings` Standard Library
	var newString string = "hello, I am Wang So!"

	var newString1 string = strings.ReplaceAll(newString, "hello", "hi") //* Replace All
	fmt.Printf("New String1: %s \n", newString1)

	var newString2 string = strings.ToUpper(newString) //* To Upper Case. There is also Lower Case.
	fmt.Printf("New String2: %s \n", newString2)

	var index int = strings.Index(newString, "I") //* If the item is not found then the index will be `-1`
	fmt.Printf("Index: %v \n", index)

	var splitString []string = strings.Split(newString, " ") //* Splits String
	fmt.Printf("Split String Slice: %v \n", splitString)

	fmt.Printf("New String: %s \n", newString)

	//* `Sort` Standard Library
	var numberSlice []int = []int{7, 4, 3, 1, 0}
	sort.Ints(numberSlice) //* sort.Ints sorts integer slices & change the original slice
	fmt.Printf("Sorted Number Slice: %v \n", numberSlice)

	var intIndex int = sort.SearchInts(numberSlice, 7)
	fmt.Printf("Number Slice Index of 7: %v \n", intIndex)

	var stringSlice []string = []string{"Wang", "So", "Shukhan", "SKN"}
	sort.Strings(stringSlice) //* sort.Strings sorts string slices & change the original slice
	fmt.Printf("Sorted String Slice: %v \n", stringSlice)

	var stringIndex int = sort.SearchStrings(stringSlice, "Wang")
	fmt.Printf("String Slice Index of Wang: %v \n", stringIndex)

	//* function as argument
	var nameSlice []string = []string{"Wang", "So", "SKN", "Shukhan"}

	utils.GetCycle(nameSlice, utils.GetGreeting)

	var fn1, ln1 string = utils.GetInitials("Wang So")
	fmt.Printf("First Name: %s & Last Name: %s \n", fn1, ln1)

	var fn2, ln2 string = utils.GetInitials("Wang So")
	fmt.Printf("First Name: %s & Last Name: %s \n", fn2, ln2)

	//}

	components.Channel()

	components.DeferPrint()

	//* Receiver Function

	var myBills *components.BillType = components.GetBill("Wang So")

	myBills.AddItem("pizza", 2, 15.5)
	myBills.AddItem("bread", 5, 12)
	myBills.AddItem("coffee", 1, 2)
	myBills.SetTip(10.0)

	fmt.Println(myBills.Format())

	var user1 utils.DataType = utils.GetData()

	fmt.Println(user1.Format())

	fmt.Println("###################")

	var newBill *components.BillType = components.CreateBill()

	fmt.Println(newBill.Format())

	fmt.Println("###################")

	wg.Wait() //* It should be at the end of "Main" thread code
	//* wg.Wait() blocks until the WaitGroup count is 0 which is set by wg.Add()
}
