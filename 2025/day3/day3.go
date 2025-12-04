package day3

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day = 3

// Return the maximum number in the given string using the
// given start and end indexes
func getMaxNumber(line string, start int, end int) (int, int) {

	if start < 0 || start > end || end > len(line) {
		panic("Invalid start or end index")
	}

	max := 0
	maxPos := 0
	for i := start; i < end; i++ {
		value := input.GetInt(line[i : i+1])

		if value > max {
			max = value
			maxPos = i
		}
	}

	return max, maxPos
}

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	total := 0

	for _, line := range lines {
		firstNumber, firstNumberPos := getMaxNumber(line, 0, len(line)-1)
		secondNumber, _ := getMaxNumber(line, firstNumberPos+1, len(line))

		total += input.GetInt(fmt.Sprintf("%d%d", firstNumber, secondNumber))
	}

	fmt.Printf("Part 1 answer: %d\n", total)
}

func Part2() {
	// @note: could use this method for part 1!
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	numberLength := 12 // desired number of digits

	total := 0

	for _, line := range lines {
		start := 0

		numberStr := ""

		for i := range numberLength {
			remaining := numberLength - i - 1
			// work out where we can search up to
			end := len(line) - remaining
			maxNum, maxNumPos := getMaxNumber(line, start, end)
			// update new start position
			start = maxNumPos + 1
			numberStr += fmt.Sprintf("%d", maxNum)
		}

		number := input.GetInt(numberStr)
		total += number
	}

	fmt.Printf("Part 2 answer: %d\n", total)
}
