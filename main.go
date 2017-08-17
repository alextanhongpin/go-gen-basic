package main

import (
	"fmt"
)

//go:generate echo Hello, Go Generate!

var VersionString = "unset"

func main() {
	fmt.Println("Version:", VersionString)
}
