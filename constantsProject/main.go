package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	const (
		feetInMeters float64 = 0.3048
		feetInYards  float64 = feetInMeters / 0.9144
	)
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := feet * feetInYards

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
