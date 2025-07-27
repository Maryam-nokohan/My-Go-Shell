package commands

import (
	"os"
	"os/exec"
	"strings"
)

func ExecuteInput(input string) error {

	input = strings.TrimSuffix(input , "\n")
	
	cmd := exec.Command(input)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}