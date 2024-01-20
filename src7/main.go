package main

//* Hello!
import "fmt"

// * `defer` keyword practice
func main() {
	defer fmt.Println("One") //* The first defer will execute last
	defer fmt.Println("Two")
	defer fmt.Println("Three") //* The last defer will execute first
	fmt.Println("Hello SKN!")

	getDefer()

	//* Expected output
	/*
		Hello SKN!
		4
		3
		2
		1
		0
		Three
		Two
		One
	*/
}

func getDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}
}
