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

func checkRepeats(leftId int, rightId int) (repeats int) {
	// 105 106... .. 110, 111
	var count int = 0
	for i := leftId; i <= rightId; i++ {
		id := strconv.Itoa(i)
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

func GiftShop() (invalidIds int, err error) {
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
		count := checkRepeats(leftId, rightId)
		totalSum += count
	}

	testCase := checkRepeats(11, 22)
	fmt.Printf("Sum of all repeats: %d\n", totalSum)
	fmt.Printf("another smaller test case: %d\n ", testCase)
	return 0, err
}

// 8123221734-8123333968

// 2665-4538
// 189952-274622
// 4975-9031
