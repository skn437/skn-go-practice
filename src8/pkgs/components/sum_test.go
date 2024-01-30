package components

import "testing"

func TestGetSum(test *testing.T) {
	var expect int = GetSum(1, 2, 3)

	if expect != 6 {
		test.Errorf("Error! \n")
	}
}
