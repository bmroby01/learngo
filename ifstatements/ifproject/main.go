package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: [username]  [password]")
		return
	}

	username := args[1]
	password := args[2]

	if username == "brianna" && password == "1234" {
		fmt.Printf("Access granted to %q.\n", username)
	} else if username == "brianna" && password != "1234" {
		fmt.Printf("Invalid password for %q.\n", username)
	} else if username == "jack" && password == "3456" {
		fmt.Printf("Access granted to %q.\n", username)
	} else if username == "jack" && password != "3456" {
		fmt.Printf("Invalid password for %q.\n", username)
	} else {
		fmt.Printf("Access denied for %q.\n", username)
	}

	//Solution from course

	//u, p := args[1], args[2]
	//if userName != "jack" {
	//	fmt.Printf("Access denied for %q.\n", u
	//} else if p != "1888" {
	//	fmt.Printf("Invalid password for %q.\n", u)
	//} else {
	//	fmt.Printf("Access granted to %q.\n", u)
	//}
}
