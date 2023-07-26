package main

import "fmt"

func main() {
	// Booleans
	myName := "Micah"

	fmt.Println("What is your first name?")
	var userName string
	fmt.Scanln(&userName)

	// If/Else Statements
	if userName == myName {
		fmt.Println("We have the same name!")
	} else {
		fmt.Println("We have different names!")
	}
}
