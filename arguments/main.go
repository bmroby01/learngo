package main

import (
	"fmt"
	"os"
)

func main() {
	// EXERCISE: Count the Arguments
	//
	//  Print the count of the command-line arguments
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 names.
	// ---------------------------------------------------------

	// UNCOMMENT & FIX THIS CODE
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)
}
