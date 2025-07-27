package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	userName := ""
	fmt.Print("Hello dear, welcome to your go SHELL! please write a username for yourself other wise you see \">$\" boring symbol all the time you can change it by writing username \"Your_UserName\" \n")
	for{
	fmt.Print(userName , ">$")
	input , err := GetInput()
	input = strings.TrimSuffix(input , "\n")
	words := strings.Split(input , " ")
	if words[0] == "username"{
		userName = words[1]
		continue
	}
	if err != nil {
		fmt.Fprintln(os.Stderr , err)
	}

	if err = commands.ExecuteInput(input); err != nil {
		fmt.Fprintln(os.Stderr , err)
	}
	
}

}