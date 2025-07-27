package commands

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ExecuteInput(input string) error {


	args := strings.Split(input , " ")
	
	
	switch args[0] {
	case "cd" :
		if len(args) < 2 {
			return errors.New("path requierd!")
		}
		return os.Chdir(args[1])
	case "exit" :
		os.Exit(0)	
	}
	
	
	cmd := exec.Command(args[0] , args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}