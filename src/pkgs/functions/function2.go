package functions

import "fmt"

func GetMultipleReturns(name string) (uint8, uint8, uint8) { //* Go can return multiple values at the same time
	var myName string = name

	var num1, num2, num3 uint8 = 1, 2, 3 //* Go can declare multiple variables at the same time

	fmt.Printf("You name is: %q\n", myName)

	return num1, num2, num3
}
