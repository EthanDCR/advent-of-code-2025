package main

import (
	"fmt"

	"example.com/adventofcode/puzzles"
)

func main() {
	giftShop, err := puzzles.GiftShop()
	if err != nil {
		panic("error!")
	}

	fmt.Println(giftShop)
}
