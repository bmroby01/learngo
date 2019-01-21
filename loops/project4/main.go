package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game!
	 
	 The program will pick %d random numbers.
	 Your mission is to guess one of those numbers.
	 
	 The greater your number is, the harder it gets.
	 
	 Wanna play?`
)

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	var guess2 int
	if len(args) == 2 {
		guess2, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
	}

	if guess < 0 || guess2 < 0 {
		fmt.Println("Please pick positive numbers")
	}

	min := guess
	if guess < guess2 {
		min = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(min + 1)

		if n == guess || n == guess2 {
			fmt.Println("YOU WIN!!")
			return
		}
	}
	fmt.Println("BAD LUCK...Try again?")
}
