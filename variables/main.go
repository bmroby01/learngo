package main

import "fmt"

func main() {
	//Print a few integer literals
	fmt.Println(
		-200, -100, 0, 50, 100,
	)

	//Print a few float literals
	fmt.Println(
		-50.5, -20.5, -0., -2.3,
	)

	//Print a few bool literals
	fmt.Println(
		true, false,
	)

	//Print your name using a string literal
	fmt.Println("Brianna")

	//Print an empty string literal, print a non-english sentence using a string literal
	fmt.Println(
		"", //empty string
		"hi",
		"مرحبًا ، اسمي براينا.", //sentence in Arabic
	)

	//Print 10 in hexadecimal
	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)

	// ---------------------------------------------------------
	// EXERCISE: Make It Blue
	//
	//  1. Change `color` variable's value to "blue"
	//
	//  2. Print it
	//
	// EXPECTED OUTPUT
	//  blue
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	color := "green"

	// ADD YOUR CODE BELOW:

	color = "blue"

	fmt.Println(color)

	// EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//
	//  2. Do not assign "dark green" to `color` directly
	//
	//     Instead, use the `color` variable again
	//     while assigning to `color`
	//
	//  3. Print it
	//
	// RESTRICTIONS
	//  WRONG ANSWER, DO NOT DO THIS:
	//  `color = "dark green"`
	//
	// HINT
	//  + operator can concatenate string values
	//
	//  Example:
	//
	//  "h" + "e" + "y" returns "hey"
	//
	// EXPECTED OUTPUT
	//  dark green
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	color = "green"

	// ADD YOUR CODE BELOW

	color = "dark " + color

	// UNCOMMENT THE CODE BELOW TO PRINT THE VARIABLE

	fmt.Println(color)
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//
	//  2. Print the `n` variable
	//
	// HINT
	//  Example: 3 * 2 = 6
	//  * is the product operator (it multiplies numbers)
	//
	// EXPECTED OUTPUT
	//  6.28
	// ---------------------------------------------------------

	// DON'T TOUCH THIS

	// Declares a new float64 variable
	// 0. means 0.0
	n := 0.

	n = 2 * 3.14

	// ?

	fmt.Println(n)

	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//
	//  2. Assign the result to the `perimeter` variable
	//
	//  3. Print the `perimeter` variable
	//
	// HINT
	//  Formula = 2 times the width and height
	//
	// EXPECTED OUTPUT
	//  22
	//
	// BONUS
	//  Find more formulas here and calculate them in new programs
	//  https://www.mathsisfun.com/area.html
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	var (
		perimeter     int
		width, height = 5, 6
	)

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT

	// ADD YOUR CODE BELOW

	perimeter = 2 * (width + height)

	fmt.Println(perimeter)

	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  go version 2
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS
	var (
		lang    string
		version int
	)

	// ADD YOUR CODE BELOW

	lang, version = "go", 2

	// DO NOT TOUCH THIS
	fmt.Println(lang, "version", version)

	// EXERCISE: Multi Assign #2
	//
	//  1. Assign the correct values to the variables
	//     to match to the EXPECTED OUTPUT below
	//
	//  2. Print the variables
	//
	// HINT
	//  Use multiple Println calls to print each sentence.
	//
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on " + planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	// EXERCISE: Multi Short Func
	//
	// 	1. Declare two variables using multiple short declaration syntax
	//
	//  2. Initialize the variables using `multi` function below
	//
	//  3. Discard the 1st variable's value in the declaration
	//
	//  4. Print only the 2nd variable
	//
	// NOTE
	//  You should use `multi` function
	//  while initializing the variables
	//
	// EXPECTED OUTPUT
	//  4
	// ---------------------------------------------------------

	// ADD YOUR DECLARATIONS HERE
	//
	_, b := multi()

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)

	// EXERCISE: Swapper
	//
	//  1. Change `color` to "orange"
	//     and `color2` to "green" at the same time
	//
	//     (use multiple-assignment)
	//
	//  2. Print the variables
	//
	// EXPECTED OUTPUT
	//  orange green
	// ---------------------------------------------------------

	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)

	// ?
}

// multi is a function that returns multiple int values

func multi() (int, int) {
	return 5, 4
}
