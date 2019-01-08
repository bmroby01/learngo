package newpackage

import (
	"fmt"
	"runtime"
)

//Version is a function to show the current version of Go
func Version() {
	fmt.Println(runtime.Version())
}
