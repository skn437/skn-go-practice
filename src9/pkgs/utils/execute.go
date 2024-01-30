package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func GetCommandExecution(name string, args ...string) {
	var command *exec.Cmd = exec.Command(name, args...)

	output, err := command.Output()

	if err != nil {
		log.Fatalf("Couldn't Execute The Command")
		return
	}

	fmt.Println(string(output))
}
