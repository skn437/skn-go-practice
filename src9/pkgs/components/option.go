package components

import (
	"fmt"
	"skn-go-practice/src9/pkgs/utils"
)

func GetOptions() {
	var options string = `
Choose a platform [e.g. Type 'a' for Linux]:
(a) Linux
(b) Project
(c) Web Services

Your Selection`

	var selection string = utils.GetReaderPrompt(options)

	if selection == "a" {
		GetLinuxOption()
	} else {
		fmt.Printf("Oops! \n")
	}
}
