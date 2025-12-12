package puzzles

// The ranges are separated by commas (,); each range gives
// its first ID and last ID separated by a dash (-).
// first lets grab the input file and parse through it.

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkDouble(id string) (value bool, err error) {
	length := len(id)
	if length == 0 {
		return false, err
	}

	// Iterate through possible pattern lengths.
	// A pattern can be at most half the length of the string.
	for i := 1; i <= length/2; i++ {
		// If the string length is not divisible by the current pattern length,
		// it cannot be a repetition of this pattern.
		if length%i != 0 {
			continue
		}

		// Extract the potential pattern.
		pattern := id[:i]

		// Construct the expected repeated string.
		// Use strings.Repeat to efficiently create the repeated string.
		repeated := strings.Repeat(pattern, length/i)

		// Compare the original string with the constructed repeated string.
		if id == repeated {
			return true, err
		}
	}

	return false, err // No repeating pattern found.
}

func checkRepeatsPartTwo(leftId int, rightId int) (repeats int) {
	// 105 106... .. 110, 111
	var count int = 0
	for i := leftId; i <= rightId; i++ {
		id := strconv.Itoa(i)
		isRepeated, err := checkDouble(id)
		if err != nil {
			panic("error calling checkDouble")
		}
		if isRepeated {
			tempId, err := strconv.Atoi(id)
			if err != nil {
				panic("error1")
			}
			count += tempId
			continue
		}

		length := len(id)
		if length%2 != 0 {
			continue
		}

		half := length / 2
		if id[:half] == id[half:] {
			temp, err := strconv.Atoi(id)
			if err != nil {
				panic("error turning temp var into int")
			}
			count += temp
		}

	}
	return count
}

func GiftShopPartTwo() (invalidIds int, err error) {
	data, err := os.ReadFile("public/day2Input.txt")
	if err != nil {
		panic("error reading the file")
	}
	inputFileString := string(data)

	lines := strings.Split(inputFileString, ",")

	var totalSum int = 0
	for _, id := range lines[:len(lines)-1] {
		splitId := strings.Split(id, "-")

		leftString := strings.TrimSpace(splitId[0])
		rightString := strings.TrimSpace(splitId[1])

		leftId, err := strconv.Atoi(leftString)
		if err != nil {
			panic("error converting left side into int")
		}
		rightId, err := strconv.Atoi(rightString)
		if err != nil {
			panic("error converting right side into int")
		}
		//		fmt.Printf("%d <- Left | right -> %d\n", leftId, rightId)
		count := checkRepeatsPartTwo(leftId, rightId)
		totalSum += count
	}

	fmt.Printf("Sum of all repeats: %d\n", totalSum)
	checkDoubleTest, err := checkDouble("23123123")
	fmt.Println(checkDoubleTest)
	return 0, err
}

// 8123221734-8123333968

// 2665-4538
// 189952-274622
// 4975-9031
