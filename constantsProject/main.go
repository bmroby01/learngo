package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)

	const m = 0.3048
	const f = 0.3333

	meters := feet * m
	yards := feet * f

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
