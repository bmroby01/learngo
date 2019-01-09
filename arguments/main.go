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

	// EXERCISE: Print the Path
	//
	//  Print the path of the running program by getting it
	//  from `os.Args` variable.
	//
	// HINT
	//  Use `go build` to build your program.
	//  Then run it using the compiled executable program file.
	//
	// EXPECTED OUTPUT SHOULD INCLUDE THIS
	//  myprogram
	// ---------------------------------------------------------

	fmt.Println(os.Args[0])

	// EXERCISE: Print Your Name
	//
	//  Get it from the first command-line argument
	//
	// INPUT
	//  Call the program using your name
	//
	// EXPECTED OUTPUT
	//  It should print your name
	//
	// EXAMPLE
	//  go run main.go inanc
	//
	//    inanc
	//
	// BONUS: Make the output like this:
	//
	//  go run main.go inanc
	//    Hi inanc
	//    How are you?
	// ---------------------------------------------------------

	// get your name from the command-line
	// and print it

	// name := os.Args[1]

	// fmt.Println("Hi", name)
	// fmt.Println("How are you?")

	// EXERCISE: Greet More People
	//
	// RESTRICTIONS
	//  1. Be sure to match to the expected output below
	//  2. Store the length of the arguments in a variable
	//  3. Store the all the arguments in variables as well
	//
	// INPUT
	//  bilbo balbo bungo
	//
	// EXPECTED OUTPUT
	//  There are 3 people!
	//  Hello great bilbo !
	//  Hello great balbo !
	//  Hello great bungo !
	//  Nice to meet you all.
	// ---------------------------------------------------------

	// TYPE YOUR CODE HERE

	// name1 := os.Args[1]
	// name2 := os.Args[2]
	// name3 := os.Args[3]

	// fmt.Println("Hello great", name1, "!")
	// fmt.Println("Hello great", name2, "!")
	// fmt.Println("Hello great", name3, "!")

	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.

	// EXERCISE: Greet 5 People
	//
	//  Greet 5 people this time.
	//
	//  Please do not copy paste from the previous exercise!
	//
	// RESTRICTION
	//  This time do not use variables.
	//
	// INPUT
	//  bilbo balbo bungo gandalf legolas
	//
	// EXPECTED OUTPUT
	//  There are 5 people!
	//  Hello great bilbo !
	//  Hello great balbo !
	//  Hello great bungo !
	//  Hello great gandalf !
	//  Hello great legolas !
	//  Nice to meet you all.
	// ---------------------------------------------------------

	fmt.Println("There are", len(os.Args)-1, "people!")
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hello great", os.Args[2], "!")
	fmt.Println("Hello great", os.Args[3], "!")
	fmt.Println("Hello great", os.Args[4], "!")
	fmt.Println("Hello great", os.Args[5], "!")
}
