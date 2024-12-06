package day3

import (
	"aoc/input"
	"fmt"
	"regexp"
)

const day int = 3

func getTotal(contents string) int {
	mulRe := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	numRe := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	var total int = 0

	for _, mulMatch := range mulRe.FindAllString(contents, -1) {
		matches := numRe.FindStringSubmatch(mulMatch)
		x := input.GetInt(matches[1])
		y := input.GetInt(matches[2])

		total += x * y
	}

	return total
}

func Part1() {
	contents := input.GetFileContents(day)
	total := getTotal(contents)

	fmt.Printf("Part 1 answer: %d\n", total)
}

func Part2() {
	contents := input.GetFileContents(day)

	contentsLen := len(contents)
	newContents := ""

	i := 0
	discard := false

	// loop through string and remove don't()/do() blocks
	for {
		if i+7 < contentsLen-1 && contents[i:i+7] == "don't()" {
			discard = true
		} else if i+4 < contentsLen-1 && contents[i:i+4] == "do()" {
			discard = false
			// shift our pointer
			i += 4
		}

		if !discard {
			newContents += contents[i : i+1]
		}

		i += 1

		if i > contentsLen-1 {
			break
		}
	}

	total := getTotal(newContents)

	fmt.Printf("Part 2 answer: %d\n", total)
}
