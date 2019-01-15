package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// STORY
	//
	//  Your boss wants you to create a program that will check
	//  whether a person can watch a particular movie or not.
	//
	//  Imagine that another program provides the age to your
	//  program. Depending on what you return, the other program
	//  will issue the tickets to the person automatically.
	//
	// EXERCISE: Movie Ratings
	//
	//  1. Get the age from the command-line.
	//
	//  2. Return one of the following messages if the age is:
	//     -> Above 17         : "R-Rated"
	//     -> Between 13 and 17: "PG-13"
	//     -> Below 13         : "PG-Rated"
	//
	// RESTRICTIONS:
	//  1. If age data is wrong or absent let the user know.
	//  2. Do not accept negative age.
	//
	// BONUS:
	//  1. BONUS: Use if statements only twice throughout your program.
	//  2. BONUS: Use an if statement only once.
	//
	// EXPECTED OUTPUT
	//  go run main.go 18
	//    R-Rated
	//
	//  go run main.go 17
	//    PG-13
	//
	//  go run main.go 12
	//    PG-Rated
	//
	//  go run main.go
	//    Requires age
	//
	//  go run main.go -5
	//    Wrong age: "-5"
	// ---------------------------------------------------------

	if len(os.Args) != 2 {
		fmt.Println("Requires age")
		return
	}

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
		return
	} else if age > 17 {
		println("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else if age < 13 {
		fmt.Println("PG-Rated")
	}

	// EXERCISE: Odd or Even
	//
	//  1. Get a number from the command-line.
	//
	//  2. Find whether the number is odd, even and divisible by 8.
	//
	// RESTRICTION
	//  Handle the error. If the number is not a valid number,
	//  or it's not provided, let the user know.
	//
	// EXPECTED OUTPUT
	//  go run main.go 16
	//    16 is an even number and it's divisible by 8
	//
	//  go run main.go 4
	//    4 is an even number
	//
	//  go run main.go 3
	//    3 is an odd number
	//
	//  go run main.go
	//    Pick a number
	//
	//  go run main.go ABC
	//    "ABC" is not a number
	// ---------------------------------------------------------

}
