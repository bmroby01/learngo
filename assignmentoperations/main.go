package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

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

	// EXERCISE: Incdecs
	//
	//  1. Increase the `counter` 5 times
	//  2. Decrease the `factor` 2 times
	//  3. Print the product of counter and factor
	//
	// RESTRICTION
	//  Use only the incdec statements
	//
	// EXPECTED OUTPUT
	//  -75
	// ---------------------------------------------------------

	// DO NOT TOUCH THIS
	counter, factor := 45, 0.5

	// TYPE YOUR CODE BELOW
	// ...

	factor--
	factor--

	counter++
	counter++
	counter++
	counter++
	counter++
	add := float64(counter) * factor
	fmt.Println(add)

	// EXERCISE: Manipulate a Counter
	//
	//  1. Write the simplest line of code to increase
	//     the counter variable by 1.
	//
	//  2. Write the simplest line of code to decrease
	//     the counter variable by 1.
	//
	//  3. Write the simplest line of code to increase
	//     the counter variable by 5.
	//
	//  4. Write the simplest line of code to multiply
	//     the counter variable by 10.
	//
	//  5. Write the simplest line of code to divide
	//     the counter variable by 5.
	//
	// EXPECTED OUTPUT
	//  10
	// ---------------------------------------------------------

	// DO NOT CHANGE THE CODE BELOW
	var counter2 int

	// TYPE YOUR CODE HERE

	counter2++

	counter2--

	counter2 += 5

	counter2 *= 10

	counter2 /= 5

	// DO NOT CHANGE THE CODE BELOW
	fmt.Println(counter2)

	// EXERCISE: Simplify the Assignments
	//
	//  Simplify the code (refactor)
	//
	// RESTRICTION
	//  Use only the incdec and assignment operations
	//
	// EXPECTED OUTPUT
	//  3
	// ---------------------------------------------------------

	width, height := 10, 2

	width++
	width += height
	width--
	width -= height
	width *= 20
	width /= 25
	width %= 5

	fmt.Println(width)

	// EXERCISE: Circle Area
	//
	//  Calculate the area of a circle from the given radius
	//
	// CIRCLE AREA FORMULA
	//  area = πr²
	//  https://en.wikipedia.org/wiki/Area#Circles
	//
	// HINT
	//  For PI you can use `math.Pi`
	//
	// EXPECTED OUTPUT
	//  radius: 10 -> area: 314.1592653589793
	//
	// BONUS EXERCISE!
	//  1. Print the area as 314.16
	//  2. To do that you need to use the correct Printf verb :)
	//      Instead of `%g` verb below.
	//
	//    EXPECTED OUTPUT
	//     radius: 10 -> area: 314.16
	// ---------------------------------------------------------

	var (
		radius = 10.
		area   float64
	)

	// ADD YOUR CODE HERE

	area = math.Pi * radius * radius

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

	// EXERCISE: Sphere Area
	//
	//  1. Get the radius from the command-line
	//  2. Convert it to a float64
	//  3. Calculate the surface area of a sphere
	//
	// SPHERE SURFACE AREA FORMULA
	//  area = 4πr²
	//  https://en.wikipedia.org/wiki/Sphere#Surface_area
	//
	// RESTRICTION
	//  Use `math.Pow` to calculate the area
	//  Read its documentation to see how it works.
	//  https://golang.org/pkg/math/#Pow
	//
	// EXPECTED OUTPUT
	//  go run main.go 10
	//    1256.64
	//
	//  go run main.go 54.2
	//    36915.47
	// ---------------------------------------------------------

	var radius1, area1 float64

	// ADD YOUR CODE HERE

	radius1, _ = strconv.ParseFloat(os.Args[1], 64)

	area1 = 4 * math.Pi * math.Pow(radius1, 2)

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %.2f\n", radius1, area1)
}
