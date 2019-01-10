package main

import (
	"fmt"
	"os"
)

func main() {
	// EXERCISE: Print Your Age
	//
	//  Print your age using Prinft
	//
	// EXPECTED OUTPUT
	//  I'm 30 years old.
	//
	// NOTE
	//  You should change 30 to your age, of course.
	// ---------------------------------------------------------

	// ?

	age := 33

	fmt.Printf("I'm %d years old.\n", age)

	// EXERCISE: Print Your Name and LastName
	//
	//  Print your name and lastname using Printf
	//
	// EXPECTED OUTPUT
	//  My name is Inanc and my lastname is Gumus.
	//
	// BONUS
	//  Store the formatting specifier (first argument of Printf)
	//    in a variable.
	//  Then pass it to printf
	// ---------------------------------------------------------

	// BONUS: Use a variable for the format specifier

	firstName := "Brianna"
	lastName := "Roby"

	fmt.Printf("My name is %s and my last name is %s.\n", firstName, lastName)

	msg := "My name is %s and my last name is %s. \n"
	fmt.Printf(msg, "Brianna", "Roby")

	// EXERCISE: False Claims
	//
	//  Use printf to print the expected output using a variable.
	//
	// EXPECTED OUTPUT
	//  These are false claims.
	// ---------------------------------------------------------

	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE

	fmt.Printf("These are %t claims.\n", tf)

	// ?

	// EXERCISE: Print the Temperature
	//
	//  Print the current temperature in your area using Printf
	//
	// NOTE
	//  Do not use %v verb
	//  Output "shouldn't" be like 29.500000
	//
	// EXPECTED OUTPUT
	//  Temperature is 29.5 degrees.
	// ---------------------------------------------------------

	fmt.Printf("Temperature is %.1f degrees.\n", 24.2)
	// ?

	// EXERCISE: Double Quotes
	//
	//  Print "hello world" with double-quotes using Printf
	//  (As you see here)
	//
	// NOTE
	//  Output "shouldn't" be just: hello world
	//
	// EXPECTED OUTPUT
	//  "hello world"
	// ---------------------------------------------------------

	fmt.Printf("\"Hello World!\"\n")
	// ?

	// EXERCISE: Print the Type
	//
	//  Print the type and value of 3 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3 is int
	// ---------------------------------------------------------

	fmt.Printf("The type of %d is %[1]T \n", 3)
	// ?

	// EXERCISE: Print the Type #2
	//
	//  Print the type and value of 3.14 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3.14 is float64
	// ---------------------------------------------------------

	fmt.Printf("The type of %.2f is %[1]T \n", 3.14)
	// ?

	// EXERCISE: Print the Type #3
	//
	//  Print the type and value of "hello" using fmt.Printf
	//
	// EXPECTED OUTPUT
	// 	Type of hello is string
	// ---------------------------------------------------------

	fmt.Printf("The type of %s is %[1]T \n", "hello")
	// ?

	// EXERCISE: Print the Type #4
	//  Print the type and value of true using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of true is bool
	// ---------------------------------------------------------

	fmt.Printf("The type of %t is %[1]T \n", true)
	// ?

	// EXERCISE: Print Your Fullname
	//
	//  1. Get your name and lastname from the command-line
	//  2. Print them using Printf
	//
	// EXAMPLE INPUT
	//  Inanc Gumus
	//
	// EXPECTED OUTPUT
	//  Your name is Inanc and your lastname is Gumus.
	// ---------------------------------------------------------

	firstname, lastname := os.Args[1], os.Args[2]
	msg2 := "Your first name is %s and your last name is %s.\n"

	fmt.Printf(msg2, firstname, lastname)
	// BONUS: Use a variable for the format specifier
}
