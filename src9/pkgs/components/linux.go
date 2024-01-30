package components

import (
	"fmt"
	"skn-go-practice/src9/pkgs/utils"
)

func GetLinuxOption() {
	var options string = `
### Linux ###	
Which command would you like to run?
(a) Port
(b) Create Random ID (32 characters)
(c) Encrypt A String Into Base64
(d) Decrypt A Base64 Into String

(z) Go Back

(x) Exit

Your Selection`

	var selection string = utils.GetReaderPrompt(options)

	switch selection {
	case "a":
		fmt.Printf("Port")
	case "b":
		fmt.Printf("Random ID")
	case "c":
		fmt.Printf("Encrypt")
	case "d":
		fmt.Printf("Decrypt")
	case "z":
		utils.GetCommandExecution("bash", "./scripts/clear.sh")
		GetOptions()
	case "x":
		utils.GetCommandExecution("bash", "./scripts/exit.sh")
	default:
		utils.GetCommandExecution("bash", "./scripts/clear.sh")
		fmt.Printf("Invalid Selection! \n")
		GetLinuxOption()
	}
}
