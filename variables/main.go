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
}
