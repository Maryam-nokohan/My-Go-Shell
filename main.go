package main

import (
	"bufio"
	"fmt"
	"os"

	commands "github.com/maryam-nokohan/MyGoShell/handlers"
)
func GetInput() (string , error){

	reader := bufio.NewReader(os.Stdin)
	input , err := reader.ReadString('\n')
	if err != nil {
		return "" , err
	}
	return input , nil

}

func main(){

	for{
	fmt.Print(">")
	input , err := GetInput()
	if err != nil {
		fmt.Fprintln(os.Stderr , err)
	}

	if err = commands.ExecuteInput(input); err != nil {
		fmt.Fprintln(os.Stderr , err)
	}
	
}

}