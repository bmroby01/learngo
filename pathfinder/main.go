package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const notFound = -1

func main() {
	// Get and split the PATH environment variable

	// SplitList function automatically finds the
	// separator for the path env variable
	words := filepath.SplitList(os.Getenv("PATH"))

	// Alternative way, but above one is better:
	// words := strings.Split(
	// 	os.Getenv("PATH"),
	// 	string(os.PathListSeparator))

	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			q, w = strings.ToLower(q), strings.ToLower(w)

			if strings.Index(w, q) != notFound {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}
