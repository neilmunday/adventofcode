package day8

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day = 8

type coord struct {
	x int
	y int
}

/*
Helper method to parse input to get antenna map and grid dimensions.
*/
func getValues() (map[string][]coord, int, int) {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	antMap := make(map[string][]coord)

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				antChar := string(line[x])
				if antMap[antChar] == nil {
					antMap[antChar] = []coord{}
				}
				antMap[antChar] = append(antMap[antChar], coord{x: x, y: y})
			}
		}
	}

	return antMap, len(lines[0]), len(lines)
}

func Part1() int {

	antMap, gridWidth, gridHeight := getValues()
	antiNodeMap := make(map[coord]bool)

	for antChar := range antMap {

		for i := 0; i < len(antMap[antChar]); i++ {
			for j := i + 1; j < len(antMap[antChar]); j++ {
				x1 := antMap[antChar][i].x
				x2 := antMap[antChar][j].x
				y1 := antMap[antChar][i].y
				y2 := antMap[antChar][j].y

				xDiff := x2 - x1
				yDiff := y2 - y1

				antiNode := coord{x: x1 - xDiff, y: y1 - yDiff}

				if antiNode.x >= 0 && antiNode.x < gridWidth && antiNode.y >= 0 && antiNode.y < gridHeight {
					antiNodeMap[antiNode] = true
				}

				antiNode = coord{x: x2 + xDiff, y: y2 + yDiff}

				if antiNode.x >= 0 && antiNode.x < gridWidth && antiNode.y >= 0 && antiNode.y < gridHeight {
					antiNodeMap[antiNode] = true
				}
			}
		}
	}

	result := len(antiNodeMap)

	fmt.Printf("Part 1 answer: %d\n", result)

	return result
}

func Part2() int {

	antMap, gridWidth, gridHeight := getValues()
	antiNodeMap := make(map[coord]bool)

	for antChar := range antMap {

		for i := 0; i < len(antMap[antChar]); i++ {
			for j := i + 1; j < len(antMap[antChar]); j++ {
				x1 := antMap[antChar][i].x
				x2 := antMap[antChar][j].x
				y1 := antMap[antChar][i].y
				y2 := antMap[antChar][j].y

				antiNodeMap[coord{x: x1, y: y1}] = true
				antiNodeMap[coord{x: x2, y: y2}] = true

				xDiff := x2 - x1
				yDiff := y2 - y1

				currentX := x1
				currentY := y1

				for {
					antiNode := coord{x: currentX - xDiff, y: currentY - yDiff}

					if antiNode.x >= 0 && antiNode.x < gridWidth && antiNode.y >= 0 && antiNode.y < gridHeight {
						antiNodeMap[antiNode] = true
					} else {
						break
					}

					currentX = antiNode.x
					currentY = antiNode.y
				}

				currentX = x2
				currentY = y2

				for {
					antiNode := coord{x: currentX + xDiff, y: currentY + yDiff}

					if antiNode.x >= 0 && antiNode.x < gridWidth && antiNode.y >= 0 && antiNode.y < gridHeight {
						antiNodeMap[antiNode] = true
					} else {
						break
					}

					currentX = antiNode.x
					currentY = antiNode.y
				}
			}
		}
	}

	result := len(antiNodeMap)

	fmt.Printf("Part 2 answer: %d\n", result)

	return result
}
