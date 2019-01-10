package main

import "fmt"

func main() {
	// EXERCISE: Do Some Calculations
	//
	//  1. Print the sum of 50 and 25
	//  2. Print the difference of 50 and 15.5
	//  3. Print the product of 50 and 0.5
	//  4. Print the quotient of 50 and 0.5
	//  5. Print the remainder of 25 and 3
	//  6. Print the negation of `5 + 2`
	//
	// EXPECTED OUTPUT
	//  75
	//  34.5
	//  25
	//  100
	//  1
	//  -7
	// ---------------------------------------------------------

	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * 0.5)
	fmt.Println(50 / 0.5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))

	// EXERCISE: Fix the Float
	//
	//  Fix the program to print 2.5 instead of 2
	//
	// EXPECTED OUTPUT
	//  2.5
	// ---------------------------------------------------------

	x := float64(5) / 2
	fmt.Println(x)

	//or

	y := 5.0 / 2
	fmt.Println(y)

	// EXERCISE: Precedence
	//
	//  Change the expressions to produce the expected outputs
	//
	// RESTRICTION
	//  Use parentheses to change the precedence
	// ---------------------------------------------------------

	// This expression should print 20
	fmt.Println((10 + 5) - (5 - 10))

	// This expression should print -16
	fmt.Println((-10 + 0.5) - (1 + 5.5))

	// This expression should print -25
	fmt.Println(5 + 10*(2-5))

	// This expression should print 0.5
	fmt.Println(0.5 * (2 - 1))

	// This expression should print 24
	fmt.Println((3+1)/2*10 + 4)

	// This expression should print 15
	fmt.Println((10 / 2) * (10 % 7))

	// This expression should print 40
	// Note that you may need to use floats to solve this
	fmt.Println(100 / (5. / 2))
}
