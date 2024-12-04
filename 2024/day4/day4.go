package day4

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day int = 4

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	gridHeight := len(lines)
	gridWidth := len(lines[0])

	matches := 0

	// find horizontal matches
	for _, line := range lines {
		for i := 0; i < gridWidth; i++ {
			if i+3 < gridWidth {
				s := line[i : i+4]
				if s == "XMAS" || s == "SAMX" {
					matches += 1
				}
			}
		}
	}

	// find vertical matches
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			if j+3 < gridHeight {
				s := string(lines[j][i]) + string(lines[j+1][i]) + string(lines[j+2][i]) + string(lines[j+3][i])
				if s == "XMAS" || s == "SAMX" {
					matches += 1
				}
			}
		}
	}

	// find diagonal left to right matches
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			if i+3 < gridWidth && j+3 < gridHeight {
				s := string(lines[j][i]) + string(lines[j+1][i+1]) + string(lines[j+2][i+2]) + string(lines[j+3][i+3])
				if s == "XMAS" || s == "SAMX" {
					matches += 1
				}
			}
		}
	}

	// find diagonal right to left matches
	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			if i-3 >= 0 && j+3 < gridHeight {
				s := string(lines[j][i]) + string(lines[j+1][i-1]) + string(lines[j+2][i-2]) + string(lines[j+3][i-3])
				if s == "XMAS" || s == "SAMX" {
					matches += 1
				}
			}
		}
	}

	fmt.Printf("Part 1 answer: %d\n", matches)
}

func Part2() {
	// find X-MASes in input
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	gridHeight := len(lines)
	gridWidth := len(lines[0])

	matches := 0

	for i := 0; i < gridWidth; i++ {
		for j := 0; j < gridHeight; j++ {
			if i+2 < gridWidth && j+2 < gridHeight {
				lrStr := string(lines[j][i]) + string(lines[j+1][i+1]) + string(lines[j+2][i+2])

				if lrStr == "MAS" || lrStr == "SAM" {
					rlStr := string(lines[j+2][i]) + string(lines[j+1][i+1]) + string(lines[j][i+2])
					if rlStr == "MAS" || rlStr == "SAM" {
						matches += 1
					}
				}
			}
		}
	}

	fmt.Printf("Part 2 answer: %d\n", matches)
}
