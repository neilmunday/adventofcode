package day2

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day = 2

func Part1() {
	contents := input.GetFileContents(day)
	ranges := strings.Split(contents, ",")

	answer := 0

	for _, r := range ranges {
		x := strings.Split(r, "-")
		start := input.GetInt(x[0])
		end := input.GetInt(x[1])
		num := start

		for num <= end {

			s := fmt.Sprintf("%d", num)

			// check if we have an equal number of characters
			if len(s)%2 == 0 {
				// split the string in the middle
				middle := len(s) / 2
				left := s[:middle]
				right := s[middle:]
				// compare left and right
				if left == right {
					answer += num
				}
			}
			num++
		}
	}

	fmt.Printf("Part 1 answer: %d\n", answer)
}

func Part2() {
	contents := input.GetFileContents(day)
	ranges := strings.Split(contents, ",")

	answer := 0

	for _, r := range ranges {
		x := strings.Split(r, "-")
		start := input.GetInt(x[0])
		end := input.GetInt(x[1])
		num := start

		for num <= end {
			// turn num into string
			numStr := fmt.Sprintf("%d", num)
			numStrLen := len(numStr)

			// examine the string looking for repeated substrings
			// starting with the first char, then the first+second char etc.
			for i := range numStrLen {
				pattern := numStr[:i+1]
				// count the number of occurrences of the substring in numStr
				count := strings.Count(numStr, pattern)
				// if we have more than one occurrence and the number of
				// occurrences * the length of the substring matches the length
				// of numStr then we have a repeated pattern for the whole string
				if count > 1 && count*len(pattern) == numStrLen {
					answer += num
					break
				}
				i++
			}

			num++
		}
	}

	fmt.Printf("Part 2 answer: %d\n", answer)
}
