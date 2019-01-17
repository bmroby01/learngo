package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1]

	if len(args) != 2 {
		fmt.Println("Please enter a number")
	}

	max, err := strconv.Atoi(args)

	if err != nil {
		fmt.Println("Please enter a number")
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
