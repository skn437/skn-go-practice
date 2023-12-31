package user

import (
	"fmt"
	"skn-go-practice/src/pkgs/shared"
	"time"
)

func GetUserInput() UserType {
	var firstName string
	var lastName string
	var bookCount uint8

	fmt.Print("Enter First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter numbers of books: ")
	fmt.Scan(&bookCount)

	var user = UserType{
		firstName: firstName,
		lastName:  lastName,
		bookCount: bookCount,
	}

	return user
}

func GetUserConfirmation() {
	time.Sleep(10 * time.Second)

	var name string = "User"

	var confirmation string = fmt.Sprintf("%s Confirmed!\n", name)

	fmt.Printf("%s\n", confirmation)

	shared.Wg.Done() //* wg.Done() decreases the WaitGroup count by 1
}

func GetUserCertificate() {

	fmt.Println("Generating User Certificate...")

	time.Sleep(10 * time.Second)

	fmt.Println("User Ticket Is Ready!")

	shared.Wg.Done()
}
