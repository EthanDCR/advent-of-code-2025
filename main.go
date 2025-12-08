package main

import (
	"fmt"

	"example.com/adventofcode/puzzles"
)

func main() {
	password, err := puzzles.CrackPassword()
	fmt.Println("password: %d\n", password)
}

// start at 50
// first instruct is right 50
// count would then be 100
// check is made
// keep getting 517 which is too low
// missing checks ?
