package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// EXERCISE: Optimal Types
	//
	//  1. Choose the optimal data types for the given situations.
	//  2. Print them all
	//  3. Try converting them to lesser data types.
	//     For example try converting int64 variable to int32.
	//     Then observe the result.
	//     Search the web why the result is so?
	//
	// NOTE
	//  This is just an exercise for teaching you different
	//  data types. Do not apply it to the real-life.
	//
	//  As I said in the lectures that, premature optimization
	//  is not a good thing.
	// ---------------------------------------------------------

	// DONT FORGET: There are also unsigned data types.
	//              (For positive numbers)

	// DO NOT USE: int data type
	// Use only the ones with the bitsizes

	// ---

	// an english letter (search web for: ascii control code)
	var letter byte = 'A'
	fmt.Println("an english letter:", letter)

	// a non-english letter (search web for: unicode codepoint)
	var nonengletter rune
	nonengletter = 'Ø'
	fmt.Println("a non-english letter:", nonengletter)

	// a year in 4 digits like 2040
	var year int32 = 2040
	fmt.Println("year:", year)

	// a month in 2 digits: 1 to 12
	var month int32 = 11
	fmt.Println("month:", month)

	// the speed of the light
	var lightSpeed int64 = 670616629 //miles
	fmt.Println("the speed of the light:", lightSpeed)

	// angle of a circle
	var angle float32 = 35.8
	fmt.Println("angle of a circle:", angle)

	// PI number: 3.141592653589793
	var pi = 3.141592653589793
	fmt.Println("PI number:", pi)

	// EXERCISE: The Type Problem
	//
	//  Solve the data type problem in the program.
	//
	// EXPECTED OUTPUT
	//  width: 265 height: 265
	//  are they equal? true
	// ---------------------------------------------------------

	// FIX THIS:
	// Change the following data types to the correct
	// data types where appropriate.
	var (
		width  uint16
		height uint16
	)

	// DONT TOUCH THIS:
	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	// UNCOMMENT THIS:
	fmt.Println("are they equal?", width == height)

	// EXERCISE: Parse Arg Numbers
	//
	//  Use strconv.ParseInt function to get int8, int16, and
	//  int32, and int64 values from command-line.
	//
	// HINT
	//  The third argument to ParseInt function represents
	//  the bitsize.
	//
	//  So, giving it 8 returns an int8 convertable value;
	//  whereas 16 returns an int16 convertable value.
	//
	//  Please explore the documentation of ParseInt function
	//  and learn how it works.
	//
	// EXPECTED OUTPUT
	//   When runned like this:
	//     go run main.go 50 25000 2000000 50000000000000000 00000100
	//
	//   It should return this:
	//     int8 value is : 50
	//     int16 value is: 25000
	//     int32 value is: 2000000
	//     int64 value is: 50000000000000000
	//     00000100 is: 4
	// ---------------------------------------------------------

	// --------------------------------------
	// EXAMPLE:
	// --------------------------------------
	// How to get an int8 from command-line:
	// First argument should be a value of -128 to 127 range
	//
	// Second argument: 10 means decimal number
	// Third argument : 8 means 8-bits (int8)
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)

	// Now the val variable is int64 because ParseInt
	// returns an int64. But, since I passed 8 as its third
	// argument, it returns int8 convertable value.
	//
	// Try running the program with a value of -128 to 127
	// Running it beyond that range will result in
	// either -128 or 127.
	fmt.Println("int8 value is:", int8(val))

	// --------------------------------------
	// NOW IT'S YOUR TURN!
	// --------------------------------------

	// 1. Get an int16 value using ParseInt and convert it and print it
	val2, _ := strconv.ParseInt(os.Args[1], 10, 8)
	fmt.Println("int8 value is :", int8(val2))

	// 2. Get an int32 value using ParseInt and convert it and print it

	// 3. Get an int64 value using ParseInt and convert it and print it

	// 4. Get an int8 value using ParseInt and convert it and print it
	//    But this time, get the value in bits.
	//
	//    For example : 00000100
	//    Should print: 4
}
