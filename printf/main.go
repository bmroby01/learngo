package main

import "fmt"

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
}
