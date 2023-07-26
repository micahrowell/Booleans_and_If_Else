package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Booleans
	myName := "Micah Rowell\n"

	fmt.Println("What is your first and last name?")

	reader := bufio.NewReader(os.Stdin)
	userName, _ := reader.ReadString('\n')

	// If/Else Statements
	if userName == myName {
		fmt.Println("We have the same name!")
	} else {
		fmt.Println("We have different names!")
	}
}
