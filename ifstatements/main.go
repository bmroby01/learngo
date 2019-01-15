package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	const (
		feetInMeters = 0.3048
		feetInYards  = feetInMeters / 0.9144
	)

	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)

	// EXERCISE: Age Seasons
	//
	//  Let's start simple. Print the expected outputs,
	//  depending on the age variable.
	//
	// EXPECTED OUTPUT
	//  If age is greater than 60, print:
	//    Getting older
	//  If age is greater than 30, print:
	//    Getting wiser
	//  If age is greater than 20, print:
	//    Adulthood
	//  If age is greater than 10, print:
	//    Young blood
	//  Otherwise, print:
	//    Booting up
	// ---------------------------------------------------------

	// Change this accordingly to produce the expected outputs
	age := 20

	// Type your if statement here.

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}

	// EXERCISE: Simplify It
	//
	//  Can you simplify the if statement inside the code below?
	//
	//  When:
	//    isSphere == true and
	//    radius is equal or greater than 200
	//
	//    It will print "It's a big sphere."
	//
	//    Otherwise, it will print "I don't know."
	//
	// EXPECTED OUTPUT
	//  It's a big sphere.
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	if isSphere == true && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}

	// EXERCISE: Arg Count
	//
	//  1. Get arguments from command-line.
	//  2. Print the expected outputs below depending on the number
	//     of arguments.
	//
	// EXPECTED OUTPUT
	//  go run main.go
	//    Give me args
	//
	//  go run main.go hello
	//    There is one: "hello"
	//
	//  go run main.go hi there
	//    There are two: "hi there"
	//
	//  go run main.go i wanna be a gopher
	//    There are 5 arguments
	// ---------------------------------------------------------

	var (
		args = os.Args
		l    = len(args) - 1
	)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("there is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf("There are two: %q\n", args[2])
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}

}
