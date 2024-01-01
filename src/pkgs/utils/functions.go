package utils

import "fmt"

func GetGreeting(name string) (string, int) {
	var length int = len(name)

	fmt.Printf("Hello, %s \n", name)

	return name, length

}

func GetCycle(names []string, f func(string) (string, int)) { //* function as argument
	for _, value := range names {
		f(value)
	}
}
