package functions

func GetNewArray(number uint8) []uint8 {
	var numberArray []uint8

	numberArray = append(numberArray, 7)
	numberArray = append(numberArray, number)

	return numberArray
}
