package components

func GetSum(args ...int) int {
	var total = 0

	for _, value := range args {
		total += value
	}

	return total
}
