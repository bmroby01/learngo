package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name][positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][4]string{
		{"happy", "excited", "overjoyed", "elated"},
		{"sad", "tired", "mad", "terible"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var mi int
	if mood != "positive" {
		mi = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][n])

}
