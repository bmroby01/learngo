package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	const (
		feetInMeters = 0.3048
		feetInYards  = feetInMeters / 0.9144
	)

	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)

	if err != nil {
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
