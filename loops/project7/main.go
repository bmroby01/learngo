package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const usage = "Welcome to the Lucky Number Game"

func main() {

	turns := os.Args[2]

	maxTurns := strconv.Atoi(turns[0])

	if maxTurns < 10 {
		maxTurns = 5
	} else if maxTurns >= 10 {
		maxTurns = os.Args[2]
	}

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number")
		return
	}

	min := 10
	if guess > min {
		min = guess
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(min + 1)

		if n == guess {
			fmt.Println("YOU WIN!!")
			if turn == 1 {
				fmt.Println("You won on the first try! Great work!")
				return
			}
			return
		}
	}
	fmt.Println("BAD LUCK...Try again?")
}
