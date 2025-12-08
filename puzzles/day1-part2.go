package puzzles

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func CrackSecondPassword() (password int, err error) {
	data, err := os.ReadFile("../advent-of-code-2025/public/day1Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputFile := string(data)

	const dialSize int = 100
	var currentPosition int = 50
	var zeroCount int = 0
	lines := strings.Split(inputFile, "\n")

	for _, element := range lines {

		element = strings.TrimSpace(element)
		if len(element) < 2 {
			continue
		}
		direction := element[0]
		distanceStr := element[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			log.Fatal(err, "error turning distance into a string")
		}

		switch string(direction) {
		case "R":
			for i := 0; i < distance; i++ {
				currentPosition++
				for currentPosition >= dialSize {
					currentPosition -= dialSize
				}

				if currentPosition == 0 {
					zeroCount++
				}
			}

		case "L":
			for i := 0; i < distance; i++ {
				currentPosition--
				for currentPosition < 0 {
					currentPosition += dialSize
				}

				if currentPosition == 0 {
					zeroCount++
				}

			}
		}
	}

	return zeroCount, err
}
