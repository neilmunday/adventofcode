package day4

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day = 4

// For the given input grid the rolls that can be removed
// are found and a new grid returned with the rolls removed.
// The number of rolls removed is returned also.
func processLines(lines []string) ([]string, int) {
	width := len(lines[0])
	height := len(lines)
	total := 0

	newLines := make([]string, height)

	for y, line := range lines {
		newLines[y] = line
		for x, value := range line {
			rolls := 0
			if value == '@' {
				// look above
				if y-1 >= 0 {
					startX := x - 1
					endX := x + 1
					if startX < 0 {
						startX = 0
					}
					if x+1 >= width {
						endX = width - 1
					}
					for i := startX; i <= endX; i++ {
						if lines[y-1][i] == '@' {
							rolls++
						}
					}
				}
				// look to left
				if x-1 >= 0 && lines[y][x-1] == '@' {
					rolls++
				}
				// look to right
				if x+1 < width && lines[y][x+1] == '@' {
					rolls++
				}
				// look below
				if y+1 < height {
					startX := x - 1
					endX := x + 1
					if startX < 0 {
						startX = 0
					}
					if x+1 >= width {
						endX = width - 1
					}
					for i := startX; i <= endX; i++ {
						if lines[y+1][i] == '@' {
							rolls++
						}
					}
				}

				if rolls < 4 {
					if x == 0 {
						newLines[y] = fmt.Sprintf(".%s", newLines[y][x+1:])
					} else if x+1 < width {
						newLines[y] = fmt.Sprintf("%s.%s", newLines[y][0:x], newLines[y][x+1:])
					} else {
						newLines[y] = fmt.Sprintf("%s.", newLines[y][0:x])
					}
					total++
				}
			}
		}
	}

	return newLines, total
}

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	_, total := processLines(lines)

	fmt.Printf("Part 1 answer: %d\n", total)
}

func Part2() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	totalRemoved := 0

	// loop until no more rolls can be removed
	for {
		newLines, removed := processLines(lines)

		if removed == 0 {
			break
		}

		totalRemoved += removed
		lines = newLines
	}

	fmt.Printf("Part 2 answer: %d\n", totalRemoved)
}
