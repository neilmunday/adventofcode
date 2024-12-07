package day2

import (
	"aoc/input"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const day int = 2

func isSafe(numbers []string) bool {
	var prevVal int64 = 0

	// check diff of first two numbers
	firstNumber := input.GetInt(numbers[0])
	secondNumber := input.GetInt(numbers[1])

	increasing := false

	if firstNumber-secondNumber < 0 {
		increasing = true
	} else if firstNumber-secondNumber > 0 {
		increasing = false
	} else {
		// same, unsafe
		return false
	}

	ok := true

	for index, number := range numbers {
		n, err := strconv.ParseInt(number, 10, 32)

		if err != nil {
			panic(err)
		}

		if index == 0 {
			prevVal = n
			continue
		}

		if n == prevVal {
			// unsafe
			ok = false
			break
		}

		if increasing {
			// n > prev
			if n <= prevVal {
				// unsafe
				ok = false
				break
			}
		}

		if !increasing {
			// decreasing
			// n < prev
			if n >= prevVal {
				ok = false
				break
			}
		}

		if !(math.Abs(float64(n-prevVal)) >= 1 && math.Abs(float64(n-prevVal)) <= 3) {
			ok = false
			break
		}

		prevVal = n
	}

	return ok
}

func Part1() int {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	var safe = 0

	for _, line := range lines {
		stripped := strings.TrimSpace(line)

		numbers := strings.Fields(stripped)

		if isSafe(numbers) {
			safe += 1
		}
	}

	fmt.Printf("Part 1 answer: %d\n", safe)

	return safe
}

func Part2() int {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	var safe = 0

	for _, line := range lines {
		stripped := strings.TrimSpace(line)

		numbers := strings.Fields(stripped)

		if isSafe(numbers) {
			safe += 1
		} else {
			// would it be safe if one number was removed?
			for i := range len(numbers) {
				// copy array
				newNumbers := slices.Clone(numbers)
				// delete index from cloned array
				copy(newNumbers[i:], newNumbers[i+1:])
				newNumbers[len(newNumbers)-1] = ""
				newNumbers = newNumbers[:len(newNumbers)-1]

				if isSafe(newNumbers) {
					safe += 1
					break
				}
			}
		}
	}

	fmt.Printf("Part 2 answer: %d\n", safe)

	return safe
}
