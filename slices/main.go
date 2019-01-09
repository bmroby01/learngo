package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("Path:", os.Args[1])
	fmt.Println("Path:", os.Args[2])
	fmt.Println("Path:", os.Args[3])

	fmt.Println("Number of items inside os.Args:", len(os.Args))
}
