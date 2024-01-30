package main

import (
	"fmt"
	"os/exec"
	"skn-go-practice/src8/pkgs/utils"
	"time"
)

func main() {

	var name string = utils.GetReaderPrompt("Enter Your Name")

	fmt.Printf("Hello '%s'! \n", name)

	var title string = utils.ViperEnvVariable("SKN.Wang.So")

	fmt.Printf("My Title: %s \n", title)

	var today time.Time = time.Now()

	fmt.Printf("Today: %s \n", today.Format("2006-01-02 15:04:05 Monday"))

	//* It works from the directory of where 'go.mod' is situated
	var command *exec.Cmd = exec.Command("bash", "./scripts/hello.sh")

	output, err := command.Output()

	if err != nil {
		fmt.Printf("Command Couldn't Be Executed!: %s", err)
	}

	fmt.Println(string(output))
}
