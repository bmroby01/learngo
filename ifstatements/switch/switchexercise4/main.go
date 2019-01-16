package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// ---------------------------------------------------------
	// STORY
	//  You want to write a program that will manipulate the
	//  given strings to uppercase, lowercase, and title case.
	//
	// EXERCISE: String Manipulator
	//
	//  1. Get the operation as the first argument.
	//
	//  2. Get the string to be manipulated as the 2nd argument.
	//
	// HINT
	//  You can find the manipulation functions in the strings
	//  package Go documentation (ToLower, ToUpper, Title).
	//
	// EXPECTED OUTPUT
	//
	//  go run main.go
	//    [command] [string]
	//
	//    Available commands: lower, upper and title
	//
	//  go run main.go lower 'OMG!'
	//    omg!
	//
	//  go run main.go upper 'omg!'
	//    OMG!
	//
	//  go run main.go title "mr. charles darwin"
	//    Mr. Charles Darwin
	//
	//  go run main.go genius "mr. charles darwin"
	//    Unknown command: "genius"
	// ---------------------------------------------------------

	const usage = `Usage: [command] [string]
	Available commands: lower, upper, and title`

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	c, s := args[1], args[2]

	switch c {
	case "lower":
		fmt.Println(strings.ToLower(s))
	case "upper":
		fmt.Println(strings.ToUpper(s))
	case "title":
		fmt.Println(strings.Title(s))
	default:
		fmt.Printf("Unknown command: %q\n", c)
	}

}
