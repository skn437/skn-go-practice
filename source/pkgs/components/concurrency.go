package components

import "fmt"

func getGreeting(name string) {
	fmt.Printf("Hello, %s", name)
}

func ConcurrencyTest() {
	go getGreeting("SKN!")
	getGreeting("Wang So!")
}
