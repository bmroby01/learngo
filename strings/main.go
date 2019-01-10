package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	// ---------------------------------------------------------
	// EXERCISE: Windows Path
	//
	//  1. Change the following program

	//  path := "c:\\program files\\duper super\\fun.txt\n" +
	//  "c:\\program files\\really\\funny.png"

	//  2. It should use a raw string literal instead
	//
	// HINT
	//  Run this program first to see its output.
	//  Then you can easily understand what it does.
	//
	// EXPECTED OUTPUT
	//  Your solution should output the same as this program.
	//  Only that it should use a raw string literal instead.
	// ---------------------------------------------------------
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := `c:\program files\duper super\fun.txt" +
		"c:\program files\really\funny.png`
	fmt.Println(path)

	// EXERCISE: Print JSON
	//
	//  1. Change the following program
	//  2. It should use a raw string literal instead
	//
	// HINT
	//  Run this program first to see its output.
	//  Then you can easily understand what it does.
	//
	// EXPECTED OUTPUT
	//  Your solution should output the same as this program.
	//  Only that it should use a raw string literal instead.
	// ---------------------------------------------------------

	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json :=
		`{ 
				"Items": [ 
				"Item": {
		 		"name": "Teddy Bear"
				}
			}]
		}`

	fmt.Println(json)

	// EXERCISE: Raw Concat
	//
	//  1. Initialize the name variable
	//     by getting input from the command line
	//
	//  2. Use concatenation operator to concatenate it
	//     with the raw string literal below
	//
	// NOTE
	//  You should concatenate the name variable in the correct
	//  place.
	//
	// EXPECTED OUTPUT
	//  Let's say that you run the program like this:
	//   go run main.go inanç
	//
	//  Then it should output this:
	//   hi inanç!
	//   how are you?
	// ---------------------------------------------------------

	// uncomment the code below
	// name := "and get the name from the command-line"

	// replace and concatenate the `name` variable
	// after `hi ` below

	name := os.Args[1]

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)

	// EXERCISE: Count the Chars
	//
	//  1. Change the following program to work with unicode
	//     characters.
	//
	// INPUT
	//  "İNANÇ"
	//
	// EXPECTED OUTPUT
	//  5
	// ---------------------------------------------------------

	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)

	// EXERCISE: Improved Banger
	//
	//  Change the Banger program the work with Unicode
	//  characters.
	//
	// INPUT
	//  "İNANÇ"
	//
	// EXPECTED OUTPUT
	//  İNANÇ!!!!!
	// ---------------------------------------------------------

	msg1 := os.Args[1]

	s := msg1 + strings.Repeat("!", utf8.RuneCountInString(msg1))

	fmt.Println(s)

	// EXERCISE: ToLowercase
	//
	//  1. Look at the documentation of strings package
	//  2. Find a function that changes the letters into lowercase
	//  3. Get a value from the command-line
	//  4. Print the given value in lowercase letters
	//
	// HINT
	//  Check out the strings package from Go online documentation.
	//  You will find the lowercase function there.
	//
	// INPUT
	//  "SHEPARD"
	//
	// EXPECTED OUTPUT
	//  shepard
	// ---------------------------------------------------------

	value := os.Args[1]

	msg2 := strings.ToLower(value)

	fmt.Println(msg2)
}
