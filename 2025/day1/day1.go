package day1

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day int = 1

func Part1() {

	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	combo := 50
	zeros := 0

	for _, line := range lines {
		direction := string(line[0])
		distance := input.GetInt(line[1:])

		if direction == "L" {
			combo -= distance % 100

			if combo < 0 {
				combo = 100 - (combo * -1)
			}
		} else {

			combo += distance % 100

			if combo > 99 {
				combo -= 100
			}
		}

		if combo == 0 {
			zeros++
		}
	}

	fmt.Printf("Part 1 answer: %d\n", zeros)
}

func Part2() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	combo := 50
	zeros := 0

	for _, line := range lines {
		direction := string(line[0])
		distance := input.GetInt(line[1:])

		for range distance {
			if direction == "L" {
				combo++
			} else {
				combo--
			}

			// Check if combo divides by 100 perfectly
			if combo%100 == 0 {
				zeros++
			}
		}
	}

	fmt.Printf("Part 2 answer: %d\n", zeros)
}
